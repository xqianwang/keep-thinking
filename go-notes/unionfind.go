package go_notes

type unionFind struct {
	count int
	parents []int
}

func (this *unionFind) Find(x int) int {
	if this.parents[x] == x {
		return x
	}
	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x, y int) {
	one := this.Find(x)
	two := this.Find(y)
	if one != two {
		this.parents[one] = two
		this.count = this.count - 1
	}
}

func Init(num int) *unionFind {
	u := &unionFind{count: num}
	u.parents = make([]int, num)
	for i := 0; i < num; i++ {
		u.parents[i] = -1
	}
	return u
}

func (this *unionFind) getGroups() int {
	return this.count
}


