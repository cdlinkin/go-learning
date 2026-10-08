package stage5generics

import (
	"testing"
	"time"
)

func TestNextWorkday(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want time.Time
	}{
		{
			name: "monday",
			date: time.Date(2026, 10, 5, 10, 30, 0, 0, time.UTC),
			want: time.Date(2026, 10, 6, 10, 30, 0, 0, time.UTC),
		},
		{
			name: "friday",
			date: time.Date(2026, 10, 9, 10, 30, 0, 0, time.UTC),
			want: time.Date(2026, 10, 12, 10, 30, 0, 0, time.UTC),
		},
		{
			name: "saturday",
			date: time.Date(2026, 10, 10, 10, 30, 0, 0, time.UTC),
			want: time.Date(2026, 10, 12, 10, 30, 0, 0, time.UTC),
		},
		{
			name: "sunday",
			date: time.Date(2026, 10, 11, 10, 30, 0, 0, time.UTC),
			want: time.Date(2026, 10, 12, 10, 30, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextWorkday(tt.date)

			if !got.Equal(tt.want) {
				t.Errorf("NextWorkday() = %v, want %v", got, tt.want)
			}
		})
	}
}
