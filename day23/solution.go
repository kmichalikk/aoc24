package day23

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Day23 struct {
	graph map[string][]string
}

func (d *Day23) Init(lines []string) {
	d.graph = make(map[string][]string)
	for _, line := range lines {
		indices := strings.Split(line, "-")
		d.graph[indices[0]] = append(d.graph[indices[0]], indices[1])
		d.graph[indices[1]] = append(d.graph[indices[1]], indices[0])
	}
}

func (d *Day23) SolveSimple() string {
	triples := 0
	for a, neighbours := range d.graph {
		if a[0] != 't' {
			continue
		}
		for i := 0; i < len(neighbours); i++ {
			b := neighbours[i]
			if b[0] == 't' && b[1] > a[1] {
				continue
			}
			for j := i + 1; j < len(neighbours); j++ {
				c := neighbours[j]
				if c[0] == 't' && (b[0] == 't' && c[1] > b[1] || c[1] > a[1]) {
					continue
				}
				for _, bn := range d.graph[b] {
					if bn == c {
						triples++
					}
				}
			}
		}
	}

	return fmt.Sprint(triples)
}

func (d *Day23) SolveAdvanced() string {
	graphMap := make(map[string]map[string]bool)
	for v, n := range d.graph {
		graphMap[v] = make(map[string]bool)
		graphMap[v][v] = true
		for _, u := range n {
			graphMap[v][u] = true
		}
	}

	names := slices.Collect(maps.Keys(graphMap))

	passwords := make([][]int, 0)
	fmt.Println("Finding 4-cliques")
	k := 4
	sel := make([]int, k)
	i := 0
	for sel[0] < len(graphMap)-k {
		for i > 0 && sel[i] >= len(graphMap)-(k-i) {
			i--
		}
		sel[i]++
		for i < k-1 {
			i++
			sel[i] = sel[i-1] + 1
		}

		isClique := true
		u, v := 0, 0
		for u < len(sel) && isClique {
			for v < len(sel) && isClique {
				isClique = graphMap[names[sel[u]]][names[sel[v]]]
				v++
			}
			u++
		}

		if isClique {
			passwords = append(passwords, slices.Clone(sel))
		}
	}

	fmt.Println("\nFound", len(passwords), "cliques, trying to extend them by another vertex")

	newPasswords := passwords
	for len(newPasswords) > 0 {
		passwords = newPasswords
		newPasswords = make([][]int, 0)
		for _, password := range passwords {
			i := slices.Max(password) + 1
			for i < len(graphMap) {
				newPassword := slices.Clone(password)
				newPassword = append(newPassword, i)

				isClique := true
				u, v := 0, 0
				for u < len(newPassword) && isClique {
					for v < len(newPassword) && isClique {
						isClique = graphMap[names[newPassword[u]]][names[newPassword[v]]]
						v++
					}
					u++
				}

				if isClique {
					newPasswords = append(newPasswords, newPassword)
				}

				i++
			}
		}
		fmt.Println("Found", len(newPasswords), "cliques (with duplicates)")
	}

	fmt.Println(passwords)

	for _, password := range passwords {
		pass := make([]string, 0)
		for _, p := range password {
			pass = append(pass, names[p])
		}
		slices.Sort(pass)
		fmt.Println(pass)
	}

	return ""
}

/*
bd,bx,du,ez,gx,jj,kr,lt,lx,oa,sq,wg,yk,zv
aa,dl,eq,fh,fy,le,mb,ob,om,qk,ri,wa,wm,xc
at,by,fa,gk,hj,mc,mj,ps,ru,sf,sn,uz,vu,yc
cx,dc,dt,fl,gb,hc,ht,ic,lj,nd,oo,pt,uk,zs
ad,dd,ed,eh,er,gq,iz,jn,mk,mw,ny,pk,rw,ws
gt,hz,mq,og,pu,rc,uw,vk,wb,ws,wz,xm,yr,zk
cp,dh,eg,fv,iv,lh,ro,tz,ue,ug,uj,wp,xi,zm
cp,cr,fm,gk,gm,jy,nw,ov,oz,rz,sr,ui,xu,zu
ag,ek,hq,ku,oj,ok,oy,pe,sw,ub,vi,wn,wx,zn
cv,dw,dx,ho,jf,kb,kd,ls,mp,nj,pu,pv,ra,tv
dz,ew,ga,ge,jo,lv,nh,nt,oc,od,qa,ti,ts,xr
al,cd,fd,hn,jt,qf,qw,qx,rg,tm,vl,xe,xl,xx
ad,ba,bi,bq,bw,dg,dm,fb,ha,hx,kg,kk,sy,ur
bo,fg,fy,gi,lo,qo,so,tj,ty,um,yf,yq,zg,zx
ap,av,dh,gu,hp,iu,iy,pb,rs,sp,uf,wy,xt,zp
cw,dy,ef,gm,iw,ji,jv,ka,ob,qv,ry,ua,wt,xz
ax,cm,dn,fe,ki,la,lw,of,sm,sx,vr,xi,yj,zz
bj,du,fi,go,li,ng,nx,ol,op,pq,pz,th,yt,yu
df,eo,gs,hl,is,jl,jz,kc,mm,mt,ru,tk,tr,za
af,au,az,dj,ee,ev,gn,hd,px,py,uo,ve,vg,zq
as,cu,cv,gg,gj,he,hr,iq,jm,mi,qt,ut,wr,yh
bs,bt,cc,el,et,hc,hi,lm,nn,qh,qn,ux,zd,zr
cf,hy,ie,lr,nk,oh,pr,ql,rh,to,vz,yv,zc,zf
ca,gf,hf,hh,ih,nb,ne,qr,sd,ta,wu,xh,yl,zs
aw,ch,em,ij,il,io,ip,jm,pp,sa,vc,xe,xw,zt
bm,cz,ev,fz,kt,ll,lm,lp,nu,rm,ss,tn,uc,ys
af,au,az,dj,ee,gn,hd,km,nx,px,py,sa,uo,zq
bf,dr,gy,hv,nm,ow,ph,rl,sg,xo,yo,ze,zj,zn
cy,ey,ft,fv,jx,pr,qp,rk,sb,ve,vn,wq,xb,ya
am,dx,ii,jd,kp,lg,me,mh,ox,rd,sc,ud,ye,yn
ao,be,ce,ds,it,jq,mc,pf,pn,rr,sy,ua,vf,vq
ah,ec,ia,ig,kv,lc,nz,ou,pl,sk,tb,up,ur,zj
di,dq,ej,he,hg,jx,kj,kz,ma,ms,rt,uz,vo,yb
cj,co,cq,da,fx,ht,iu,kh,ks,lf,lu,rp,uu,we
ak,bh,db,ep,gf,ib,kx,ky,lz,pj,qc,su,sz,xt
ab,ay,cy,hu,in,lj,mv,np,nv,or,qm,uq,ww,xn
bp,cb,fr,gl,hb,hw,kf,pa,re,sh,sj,sl,vp,wv
*/
