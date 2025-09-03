package rich

import (
	"strings"
	"testing"
)

func TestSprint(t *testing.T) {
	red := FromColorName("red")

	tests := []struct {
		name   string
		format string
		args   []any
		want   string
	}{
		{
			name:   "plain text",
			format: "Hello",
			args:   nil,
			want:   "Hello" + string(Reset),
		},
		{
			name:   "bold text",
			format: "Hello",
			args:   []any{Bold},
			want:   string(Bold) + "Hello" + string(Reset),
		},
		{
			name:   "red text",
			format: "Hello",
			args:   []any{red},
			want:   red.ToANSI() + "Hello" + string(Reset),
		},
		{
			name:   "red + bold",
			format: "Hello",
			args:   []any{red, Bold},
			want:   red.ToANSI() + string(Bold) + "Hello" + string(Reset),
		},
		{
			name:   "format with value",
			format: "Value: %d",
			args:   []any{42, Bold},
			want:   string(Bold) + "Value: 42" + string(Reset),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Sprint(tc.format, tc.args...)
			if got != tc.want {
				t.Errorf("Sprint(%q, %v) = %q, want %q", tc.format, tc.args, got, tc.want)
			}
		})
	}
}

func TestSprintf(t *testing.T) {
	green := FromColorName("green")

	tests := []struct {
		name   string
		format string
		args   []any
		want   string
	}{
		{
			name:   "plain",
			format: "Hello, %s!",
			args:   []any{"world"},
			want:   "Hello, world!" + string(Reset),
		},
		{
			name:   "green style",
			format: "Hello, %s!",
			args:   []any{"world", green},
			want:   green.ToANSI() + "Hello, world!" + string(Reset),
		},
		{
			name:   "green + bold",
			format: "Value: %d",
			args:   []any{42, green, Bold},
			want:   green.ToANSI() + string(Bold) + "Value: 42" + string(Reset),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Sprintf(tc.format, tc.args...)
			if got != tc.want {
				t.Errorf("Sprintf(%q, %v) = %q, want %q", tc.format, tc.args, got, tc.want)
			}
		})
	}
}

func TestPrintln(t *testing.T) {
	var sb strings.Builder
	old := outWriter
	defer func() { outWriter = old }()
	outWriter = &sb

	bold := Bold
	red := FromColorName("red")

	tests := []struct {
		name   string
		format string
		args   []any
		want   string
	}{
		{
			name:   "plain",
			format: "Hello",
			args:   nil,
			want:   "Hello\n" + string(Reset),
		},
		{
			name:   "bold",
			format: "Hello",
			args:   []any{bold},
			want:   string(Bold) + "Hello\n" + string(Reset),
		},
		{
			name:   "red + bold",
			format: "Hello",
			args:   []any{red, bold},
			want:   red.ToANSI() + string(Bold) + "Hello\n" + string(Reset),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sb.Reset()
			Println(tc.format, tc.args...)
			got := sb.String()
			if got != tc.want {
				t.Errorf("Println(%q, %v) = %q, want %q", tc.format, tc.args, got, tc.want)
			}
		})
	}
}

func TestPrintf(t *testing.T) {
	var sb strings.Builder
	old := outWriter
	defer func() { outWriter = old }()
	outWriter = &sb

	bold := Bold
	red := FromColorName("red")

	tests := []struct {
		name   string
		format string
		args   []any
		want   string
	}{
		{
			name:   "plain",
			format: "Hello, %s!",
			args:   []any{"world"},
			want:   "Hello, world!" + string(Reset),
		},
		{
			name:   "bold",
			format: "Hello, %s!",
			args:   []any{"world", bold},
			want:   string(Bold) + "Hello, world!" + string(Reset),
		},
		{
			name:   "red + bold",
			format: "Value: %d",
			args:   []any{42, red, bold},
			want:   red.ToANSI() + string(Bold) + "Value: 42" + string(Reset),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sb.Reset()
			Printf(tc.format, tc.args...)
			got := sb.String()
			if got != tc.want {
				t.Errorf("Printf(%q, %v) = %q, want %q", tc.format, tc.args, got, tc.want)
			}
		})
	}
}
