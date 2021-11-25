package mpm

import "github.com/go-gl/mathgl/mgl64"

func UpdateDeformationGradients(ps *Particles) {
	for i, p := range ps.Ps {
		fpNew := mgl64.Ident2()
		fpNew = fpNew.Add(p.c.Mul(dt))
		p.f = fpNew.Mul2(p.f)
		ps.Ps[i] = p
	}

	//            var Fp_new = math.float2x2(
	//                1, 0,
	//                0, 1
	//            );
	//            Fp_new += dt * p.C;
	//            Fs[i] = math.mul(Fp_new, Fs[i]);
	//
	//            ps[i] = p;
}
