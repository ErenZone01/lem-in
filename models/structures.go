package models

type Room struct {
	X, Y string
	Visited bool
	Ant string
	Path string
	Links [] string
}

func (r *Room) UpdateLinks(links string) {
	r.Links = append(r.Links, links)
}

func (r *Room) MarkAsVisited() {
	r.Visited = !r.Visited
}

type Tunnel struct {
	From, To string
}

func (r *Room) InitRoom(data []string) {
	if len(data) >= 3 {
		r.X = data[1]
		r.Y = data[2]
	}
}

func (t *Tunnel) InitTunnel(data []string) {
	if len(data) >= 2 {
		t.From = data[0]
		t.To = data[1]
	}
}
