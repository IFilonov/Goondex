package hw

import "testing"

func Test_CalculateDistance(t *testing.T) {
	tests := []struct {
		name         string
		X1           float64
		X2           float64
		Y1           float64
		Y2           float64
		wantDistance float64
	}{
		{
			name:         "#1",
			X1:           1,
			X2:           4,
			Y1:           1,
			Y2:           5,
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance, err := CalculateDistance(tt.X1, tt.Y1, tt.X2, tt.Y2); gotDistance != tt.wantDistance && err == nil {
				t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
