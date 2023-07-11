package hw

import "testing"

func TestGeom_CalculateDistance(t *testing.T) {
	tests := []struct {
		name         string
		geom         *Geom
		wantDistance float64
	}{
		{
			name:         "#1",
			geom:         New(1, 1, 4, 5),
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := tt.geom.CalcDist(); gotDistance != tt.wantDistance {
				t.Errorf("Geom.CalcDist() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
