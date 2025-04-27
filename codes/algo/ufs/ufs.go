package main

type Ufs struct {
	p         []int
	cnt       []int
	component int
}

func NewUfs(size int) *Ufs {
	p := make([]int, size)
	cnt := make([]int, size)
	for i := range p {
		p[i] = i
		cnt[i] = 1
	}
	return &Ufs{p, cnt, size}
}

func (u *Ufs) union(a, b int) {
	x, y := u.find(a), u.find(b)
	if x == y {
		return
	}
	rankX, rankY := u.cnt[x], u.cnt[y]
	if rankX <= rankY {
		u.p[x] = y
		u.cnt[y] += u.cnt[x]
	} else {
		u.p[y] = x
		u.cnt[x] += u.cnt[y]
	}
	u.component--
	return
}

func (u *Ufs) find(a int) int {
	if u.p[a] != a {
		u.p[a] = u.find(u.p[a])
	}
	return u.p[a]
}