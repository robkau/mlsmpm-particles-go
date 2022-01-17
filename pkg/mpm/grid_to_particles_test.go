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
		name    string
		args    args
		want    mgl64.Mat2
		wantDet float64
	}{
		{"cellDist X 0", args{mgl64.Vec2{1, 1}, mgl64.Vec2{0, 1}}, mgl64.Mat2{0, 0, 1, 1}, 0},
		{"cellDist Y 0", args{mgl64.Vec2{1, 1}, mgl64.Vec2{1, 0}}, mgl64.Mat2{1, 1, 0, 0}, 0},
		{"both", args{mgl64.Vec2{0.22, 0.77}, mgl64.Vec2{2, -1}}, mgl64.Mat2{0.22 * 2, 0.77 * 2, 0.22 * -1, 0.77 * -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := weightedVelocityAndCellDistToTerm(tt.args.weightedVelocity, tt.args.cellDist)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("weightedVelocityAndCellDistToTerm() = %v, want %v", got, tt.want)
			}
			gotDet := got.Det()
			if gotDet != tt.wantDet {
				t.Errorf("weightedVelocityAndCellDistToTerm().Det() = %v, want %v", gotDet, tt.wantDet)
			}

		})
	}
}
