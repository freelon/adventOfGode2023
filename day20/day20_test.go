package day20

import (
	"adventOfGode2023/util"
	"fmt"
	"strconv"
	"testing"
)

const inputA = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

func TestPart1a(t *testing.T) {
	util.Assert(t, 32000000, Part1(inputA))
}
func TestPart1b(t *testing.T) {
	util.Assert(t, 11687500, Part1("broadcaster -> a\n%a -> inv, con\n&inv -> b\n%b -> con\n&con -> output"))
}
func TestPart2(t *testing.T) {
	a, _ := strconv.Atoi(Part2(inputKF))
	b, _ := strconv.Atoi(Part2(inputKR))
	c, _ := strconv.Atoi(Part2(inputZS))
	d, _ := strconv.Atoi(Part2(inputQK))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a * b * c * d)
}

const inputKF = `%mh -> rz
%nd -> jx
%xt -> cx
%dp -> mh
%pz -> zg, bf
%rp -> jb, bf
%jb -> bf, kp
%rj -> xt, cx
%hg -> dl, bf
%pt -> gm, vv
%pf -> xk, qr
%cv -> jp, cx
%zg -> bb
%qn -> gm, bh
%kp -> pz
%kg -> gm, pt
%sl -> rp
%dz -> bf, dc
%hm -> cx, tz
%dc -> fk
%xk -> qr, sf
%bq -> qr, mg
%sf -> qr
&cx -> ff, vx, zs
%hr -> fq, gm
%ls -> lf, gm
%mf -> cx, sx
%vq -> gm
%sx -> cx, rj
&gm -> kg, kf, fq, nc, lf
%jx -> qr, zz
%tz -> mf, cx
%jp -> cx, kt
%bb -> hg, bf
%zz -> pf, qr
&qr -> dp, bq, nd, rz, mg, qk, mh
%nc -> gb
%kt -> hm, cx
%mg -> dp
%dl -> bf
&bf -> dz, zg, kr, sl, fk, kp, dc
%bh -> vq, gm
&kf -> gf
%fq -> qn
%vl -> vx, cx
%fk -> sl
%tj -> nd, qr
%gb -> ls, gm
%lf -> hr
%vx -> cv
%ff -> vl, cx
broadcaster -> kg
%vv -> nc, gm
&gf -> rx
%rz -> tj`

const inputKR = `%mh -> rz
%nd -> jx
%xt -> cx
%dp -> mh
%pz -> zg, bf
%rp -> jb, bf
%jb -> bf, kp
%rj -> xt, cx
%hg -> dl, bf
%pt -> gm, vv
%pf -> xk, qr
%cv -> jp, cx
%zg -> bb
%qn -> gm, bh
%kp -> pz
%kg -> gm, pt
%sl -> rp
%dz -> bf, dc
%hm -> cx, tz
%dc -> fk
%xk -> qr, sf
&kr -> gf
%bq -> qr, mg
%sf -> qr
&cx -> ff, vx, zs
%hr -> fq, gm
%ls -> lf, gm
%mf -> cx, sx
%vq -> gm
%sx -> cx, rj
&gm -> kg, kf, fq, nc, lf
%jx -> qr, zz
%tz -> mf, cx
%jp -> cx, kt
%bb -> hg, bf
%zz -> pf, qr
&qr -> dp, bq, nd, rz, mg, qk, mh
%nc -> gb
%kt -> hm, cx
%mg -> dp
%dl -> bf
&bf -> dz, zg, kr, sl, fk, kp, dc
%bh -> vq, gm
%fq -> qn
%vl -> vx, cx
%fk -> sl
%tj -> nd, qr
%gb -> ls, gm
%lf -> hr
%vx -> cv
%ff -> vl, cx
broadcaster -> dz
%vv -> nc, gm
&gf -> rx
%rz -> tj`

const inputZS = `%mh -> rz
%nd -> jx
%xt -> cx
%dp -> mh
%pz -> zg, bf
%rp -> jb, bf
%jb -> bf, kp
%rj -> xt, cx
%hg -> dl, bf
%pt -> gm, vv
%pf -> xk, qr
%cv -> jp, cx
%zg -> bb
%qn -> gm, bh
%kp -> pz
%kg -> gm, pt
%sl -> rp
%dz -> bf, dc
%hm -> cx, tz
%dc -> fk
%xk -> qr, sf
%bq -> qr, mg
%sf -> qr
&cx -> ff, vx, zs
%hr -> fq, gm
%ls -> lf, gm
%mf -> cx, sx
%vq -> gm
%sx -> cx, rj
&gm -> kg, kf, fq, nc, lf
%jx -> qr, zz
%tz -> mf, cx
%jp -> cx, kt
%bb -> hg, bf
%zz -> pf, qr
&qr -> dp, bq, nd, rz, mg, qk, mh
%nc -> gb
%kt -> hm, cx
%mg -> dp
%dl -> bf
&zs -> gf
&bf -> dz, zg, kr, sl, fk, kp, dc
%bh -> vq, gm
%fq -> qn
%vl -> vx, cx
%fk -> sl
%tj -> nd, qr
%gb -> ls, gm
%lf -> hr
%vx -> cv
%ff -> vl, cx
broadcaster -> ff
%vv -> nc, gm
&gf -> rx
%rz -> tj`

const inputQK = `%mh -> rz
%nd -> jx
%xt -> cx
%dp -> mh
%pz -> zg, bf
%rp -> jb, bf
%jb -> bf, kp
%rj -> xt, cx
%hg -> dl, bf
%pt -> gm, vv
%pf -> xk, qr
%cv -> jp, cx
%zg -> bb
%qn -> gm, bh
%kp -> pz
%kg -> gm, pt
%sl -> rp
%dz -> bf, dc
%hm -> cx, tz
%dc -> fk
%xk -> qr, sf
%bq -> qr, mg
%sf -> qr
&cx -> ff, vx, zs
%hr -> fq, gm
%ls -> lf, gm
%mf -> cx, sx
%vq -> gm
%sx -> cx, rj
&gm -> kg, kf, fq, nc, lf
%jx -> qr, zz
%tz -> mf, cx
%jp -> cx, kt
%bb -> hg, bf
%zz -> pf, qr
&qr -> dp, bq, nd, rz, mg, qk, mh
%nc -> gb
%kt -> hm, cx
%mg -> dp
%dl -> bf
&bf -> dz, zg, kr, sl, fk, kp, dc
%bh -> vq, gm
%fq -> qn
%vl -> vx, cx
&qk -> gf
%fk -> sl
%tj -> nd, qr
%gb -> ls, gm
%lf -> hr
%vx -> cv
%ff -> vl, cx
broadcaster -> bq
%vv -> nc, gm
&gf -> rx
%rz -> tj`
