package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"reflect"
	"testing"
)

func Test_weightedVelocityAndCellDistToTerm(t *testing.T) {
	type args struct {
		weightedVelocity mgl64.Vec2
		cellDist         mgl64.Vec2
	}
	tests := []struct {
		name string
		args args
		want mgl64.Mat2
	}{
		{"cellDist X 0", args{mgl64.Vec2{1, 1}, mgl64.Vec2{0, 1}}, mgl64.Mat2{0, 1, 0, 1}},
		{"cellDist Y 0", args{mgl64.Vec2{1, 1}, mgl64.Vec2{1, 0}}, mgl64.Mat2{1, 0, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := weightedVelocityAndCellDistToTerm(tt.args.weightedVelocity, tt.args.cellDist); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("weightedVelocityAndCellDistToTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
