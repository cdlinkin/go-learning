/*
exercise:

	First, write the tests, then the code for the ParseDuration("1h30m")
	function (a simplified parser: hours, minutes, seconds)
	At least 12 table cases, including erroneous input
*/
package duration

import (
	"testing"
	"time"
)

// ParseDuration("1h30m")
func TestParseDuration(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Duration
		wantErr bool
	}{
		{
			name:  "one hour",
			input: "1h",
			want:  time.Hour,
		},
		{
			name:  "thirty minutes",
			input: "30m",
			want:  30 * time.Minute,
		},
		{
			name:  "forty five seconds",
			input: "45s",
			want:  45 * time.Second,
		},
		{
			name:  "one hour thirty minutes",
			input: "1h30m",
			want:  90 * time.Minute,
		},
		{
			name:  "one hour thirty minutes twenty seconds",
			input: "1h30m20s",
			want:  time.Hour + 30*time.Minute + 20*time.Second,
		},
		{
			name:  "one minute thirty seconds",
			input: "1m30s",
			want:  90 * time.Second,
		},
		{
			name:  "zero seconds",
			input: "0s",
			want:  0,
		},
		{
			name:  "zero duration",
			input: "0h0m0s",
			want:  0,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
		{
			name:    "unknown unit",
			input:   "1d",
			wantErr: true,
		},
		{
			name:    "number without unit",
			input:   "123",
			wantErr: true,
		},
		{
			name:    "text instead of duration",
			input:   "abc",
			wantErr: true,
		},
		{
			name:    "negative duration",
			input:   "-1h",
			wantErr: true,
		},
		{
			name:    "invalid characters",
			input:   "1hxyz",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDuration(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("ParseDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
