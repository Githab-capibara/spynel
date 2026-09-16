package harness

import (
	"context"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/agent0ai/spynel/internal/core"
)

func TestACPStreamsPersistsAndResumesSessions(t *testing.T) {
	command, root, logPath := portableHarnessFixture(t, "acp-lifecycle")
	sessionsPath := filepath.Join(root, "sessions.json")
	config := HarnessConfig{
		Command: command, Args: []string{"--stdio", "value with spaces"}, Cwd: root,
		Model: "model-a", Effort: "high", Sandbox: "read-only", SessionsFile: sessionsPath, Version: "test",
	}
	adapter, err := NewACP(config)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := adapter.Start(ctx); err != nil {
		t.Fatal(err)
	}
	events := make(chan core.Event, 32)
	threadID, steered, err := adapter.Send(ctx, "chat", "test prompt", func(event core.Event) { events <- event })
	if err != nil || steered || threadID != "acp-session" {
		t.Fatalf("ACP Send() = %q, %t, %v", threadID, steered, err)
	}
	var streamed strings.Builder
	var final core.Event
	foundTool := false
	for !final.Done {
		select {
		case event := <-events:
			if event.Kind == core.EventDelta {
				streamed.WriteString(event.Text)
			}
			foundTool = foundTool || event.Kind == core.EventStatus && strings.Contains(event.Text, "Edit file")
			if event.Done {
				final = event
			}
		case <-ctx.Done():
			t.Fatal("timed out waiting for ACP result")
		}
	}
	if streamed.String() != "hello world" || final.Kind != core.EventFinal || final.Text != "hello world" || final.FinalText == nil || *final.FinalText != "hello world" || !foundTool {
		t.Fatalf("ACP events streamed=%q final=%#v tool=%t", streamed.String(), final, foundTool)
	}
	if err := adapter.Close(); err != nil {
		t.Fatal(err)
	}

	restarted, err := NewACP(config)
	if err != nil || restarted.ThreadID("chat") != "acp-session" {
		t.Fatalf("persisted ACP session = %q, %v", restarted.ThreadID("chat"), err)
	}
	if err := restarted.Start(ctx); err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{}, 1)
	if threadID, steered, err := restarted.Send(ctx, "chat", "continued", func(event core.Event) {
		if event.Done {
			done <- struct{}{}
		}
	}); err != nil || steered || threadID != "acp-session" {
		t.Fatalf("resumed ACP Send() = %q, %t, %v", threadID, steered, err)
	}
	select {
	case <-done:
	case <-ctx.Done():
		t.Fatal("timed out waiting for resumed ACP turn")
	}
	_ = restarted.Close()

	invocations, resumed, configured, rejected := 0, 0, 0, false
	for _, record := range readFixtureRecords(t, logPath) {
		if record.Kind == "invocation" {
			invocations++
			if record.Cwd != root || record.Executable != command || len(record.Args) != 2 || record.Args[0] != "--stdio" || record.Args[1] != "value with spaces" {
				t.Fatalf("portable ACP invocation = %#v", record)
			}
		}
		resumed = resumed + boolInt(record.Method == "session/resume")
		if record.Method == "session/set_config_option" {
			configured++
			params := string(record.Params)
			model := strings.Contains(params, `"configId":"model"`) && strings.Contains(params, `"value":"model-a"`)
			if !strings.Contains(params, `"sessionId":"acp-session"`) || !model || strings.Contains(params, `"configId":"thought"`) {
				t.Fatalf("ACP config option request = %s", record.Params)
			}
		}
		if record.Method == "permission-response" && strings.Contains(string(record.Params), `"optionId":"reject"`) {
			rejected = true
		}
	}
	if invocations != 2 || resumed != 1 || configured != 1 || !rejected {
		t.Fatalf("ACP evidence = invocations %d, resumes %d, config options %d, read-only rejected edit %t", invocations, resumed, configured, rejected)
	}
}

func TestACPCancelNotificationTerminatesPrompt(t *testing.T) {
	command, root, logPath := portableHarnessFixture(t, "acp-interrupt")
	adapter, err := NewACP(HarnessConfig{Command: command, Cwd: root, SessionsFile: filepath.Join(root, "sessions.json")})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := adapter.Start(ctx); err != nil {
		t.Fatal(err)
	}
	defer adapter.Close()
	done := make(chan core.Event, 1)
	if _, _, err := adapter.Send(ctx, "chat", "work", func(event core.Event) {
		if event.Done {
			done <- event
		}
	}); err != nil {
		t.Fatal(err)
	}
	if stopped, err := adapter.Interrupt(ctx, "chat"); err != nil || !stopped {
		t.Fatalf("ACP interrupt = %t, %v", stopped, err)
	}
	select {
	case event := <-done:
		if event.Kind != core.EventError || !strings.Contains(event.Text, "cancelled") {
			t.Fatalf("ACP cancelled terminal = %#v", event)
		}
	case <-ctx.Done():
		t.Fatal("timed out waiting for ACP cancellation")
	}
	foundCancel, cancelledPermission := false, false
	for _, record := range readFixtureRecords(t, logPath) {
		foundCancel = foundCancel || record.Method == "session/cancel"
		cancelledPermission = cancelledPermission || record.Method == "permission-response" &&
			strings.Contains(string(record.Params), `"outcome":{"outcome":"cancelled"}`)
	}
	if !foundCancel || !cancelledPermission {
		t.Fatalf("ACP cancellation evidence = cancel notification %t, late permission cancelled %t", foundCancel, cancelledPermission)
	}
}

func TestAgentZeroSessionFailureAddsSafeReachabilityAndAuthenticationGuidance(t *testing.T) {
	command, root, _ := portableHarnessFixture(t, "acp-session-error")
	adapter, err := NewACP(HarnessConfig{Name: "agent-zero", Command: command, Args: []string{"acp"}, Cwd: root})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := adapter.Start(ctx); err != nil {
		t.Fatal(err)
	}
	defer adapter.Close()
	_, _, err = adapter.Send(ctx, "chat", "work", func(core.Event) {})
	if err == nil || !strings.Contains(err.Error(), "ensure Agent Zero is running and reachable") || !strings.Contains(err.Error(), "authenticate or configure the connection through A0 CLI") || !strings.Contains(err.Error(), "Internal error") {
		t.Fatalf("Agent Zero session error = %v", err)
	}
}

func TestACPSelectsIdModelOptionAmongSameCategoryOptions(t *testing.T) {
	command, root, logPath := portableHarnessFixture(t, "acp-dual-model-options")
	adapter, err := NewACP(HarnessConfig{Name: "cline", Command: command, Cwd: root, Model: "anthropic/claude-sonnet-5", Sandbox: "read-only", SessionsFile: filepath.Join(root, "sessions.json")})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := adapter.Start(ctx); err != nil {
		t.Fatal(err)
	}
	defer adapter.Close()
	done := make(chan core.Event, 1)
	if _, _, err := adapter.Send(ctx, "chat", "work", func(event core.Event) {
		if event.Done {
			done <- event
		}
	}); err != nil {
		t.Fatal(err)
	}
	select {
	case <-done:
	case <-ctx.Done():
		t.Fatal("timed out waiting for ACP dual-option turn")
	}
	configured := 0
	modeConfigured := ""
	for _, record := range readFixtureRecords(t, logPath) {
		if record.Method != "session/set_config_option" {
			continue
		}
		configured++
		params := string(record.Params)
		if strings.Contains(params, `"configId":"provider"`) {
			t.Fatalf("model value leaked into provider option: %s", params)
		}
		if strings.Contains(params, `"configId":"mode"`) {
			if !strings.Contains(params, `"value":"plan"`) {
				t.Fatalf("read-only sandbox must select plan mode: %s", params)
			}
			modeConfigured = params
			continue
		}
		if !strings.Contains(params, `"configId":"model"`) || !strings.Contains(params, `"value":"anthropic/claude-sonnet-5"`) {
			t.Fatalf("ACP model config option request = %s", params)
		}
	}
	if configured != 2 || modeConfigured == "" {
		t.Fatalf("expected one mode and one model config write, got %d (mode %q)", configured, modeConfigured)
	}
}

func TestACPModelWithoutMatchingOptionFailsNewAndSucceedsResume(t *testing.T) {
	command, root, _ := portableHarnessFixture(t, "acp-provider-only-model")
	sessionsPath := filepath.Join(root, "sessions.json")
	config := HarnessConfig{Name: "cline", Command: command, Cwd: root, Model: "anthropic/claude-sonnet-5", Sandbox: "read-only", SessionsFile: sessionsPath}
	adapter, err := NewACP(config)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := adapter.Start(ctx); err != nil {
		t.Fatal(err)
	}
	_, _, err = adapter.Send(ctx, "chat", "work", func(core.Event) {})
	if err == nil || !strings.Contains(err.Error(), "does not expose a model config option for \"anthropic/claude-sonnet-5\"") {
		t.Fatalf("model without a matching option = %v", err)
	}
	if err := adapter.Close(); err != nil {
		t.Fatal(err)
	}
	persisted, err := NewACP(config)
	if err != nil {
		t.Fatal(err)
	}
	if persisted.ThreadID("chat") != "acp-session" {
		t.Fatalf("failed first session not persisted = %q", persisted.ThreadID("chat"))
	}
	if err := persisted.Start(ctx); err != nil {
		t.Fatal(err)
	}
	defer persisted.Close()
	done := make(chan core.Event, 1)
	if _, _, err := persisted.Send(ctx, "chat", "continued", func(event core.Event) {
		if event.Done {
			done <- event
		}
	}); err != nil {
		t.Fatalf("resume with a persisted session should not depend on option listing, got %v", err)
	}
	select {
	case <-done:
	case <-ctx.Done():
		t.Fatal("timed out waiting for resumed ACP turn")
	}
}

func TestACPSessionAuthErrorAddsAdvertisedMethodGuidance(t *testing.T) {
	command, root, _ := portableHarnessFixture(t, "acp-auth-required")
	adapter, err := NewACP(HarnessConfig{Name: "cline", Command: command, Args: []string{"acp"}, Cwd: root, Model: "anthropic/claude-sonnet-5"})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := adapter.Start(ctx); err != nil {
		t.Fatal(err)
	}
	defer adapter.Close()
	_, _, err = adapter.Send(ctx, "chat", "work", func(core.Event) {})
	if err == nil {
		t.Fatal("expected an authentication error")
	}
	message := err.Error()
	for _, fragment := range []string{"ACP error -32000", "Authentication required", "run `", " outside Spynel", "advertised methods: Sign in with Cline", "then retry"} {
		if !strings.Contains(message, fragment) {
			t.Fatalf("authentication guidance missing %q: %v", fragment, err)
		}
	}
}

func TestACPBuildsModelCatalogFromConfigOptions(t *testing.T) {
	command, root, logPath := portableHarnessFixture(t, "acp-dual-model-options")
	adapter, err := NewACP(HarnessConfig{Name: "cline", Command: command, Cwd: root, Sandbox: "read-only", SessionsFile: filepath.Join(root, "sessions.json")})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := adapter.Start(ctx); err != nil {
		t.Fatal(err)
	}
	defer adapter.Close()
	models, err := adapter.Models(ctx)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]string{"anthropic/claude-sonnet-5": "Claude Sonnet 5", "openai/gpt-5": "GPT-5"}
	if len(models) != len(expected) {
		t.Fatalf("ACP catalog size = %d: %#v", len(models), models)
	}
	for _, model := range models {
		if expected[model.ID] != model.DisplayName {
			t.Fatalf("ACP catalog entry = %#v", model)
		}
		if model.ID == "cline" || model.ID == "openai" {
			t.Fatalf("provider choice leaked into model catalog: %#v", model)
		}
	}
	invoked := 0
	for _, record := range readFixtureRecords(t, logPath) {
		if record.Method == "session/new" {
			invoked++
		}
		if record.Method == "session/prompt" {
			t.Fatalf("model catalog must not send prompts: %s", record.Params)
		}
	}
	if invoked != 1 {
		t.Fatalf("catalog session creation count = %d", invoked)
	}
	cached, err := adapter.Models(ctx)
	if err != nil || len(cached) != len(expected) {
		t.Fatalf("cached catalog = %#v, %v", cached, err)
	}
	invocations := 0
	for _, record := range readFixtureRecords(t, logPath) {
		invocations += boolInt(record.Kind == "invocation")
	}
	if invocations != 1 {
		t.Fatalf("cached lookup restarted the agent: %d invocations", invocations)
	}
}

func boolInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
