package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneCleanBottleVirusBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.28 19.72a3.145 3.145 0 0 1-.029-4.473a3.143 3.143 0 0 1 4.473.03m.921 2.221a3.143 3.143 0 0 1-3.143 3.143m-.525-8.643h1.049m-.524 0v2.357m5.5 2.619v1.049m0-.525h-2.357m1.117 3.519l-.742.74m.37-.37l-1.666-1.666m-1.698 3.277h-1.049m.525 0v-2.357m-5.5-2.618v-1.049m0 .524h2.356m-1.116-3.518l.74-.74m-.369.37l1.667 1.667m7.722-3.279l-11 11m-3.504-2.751h-4.5a3 3 0 0 1-3-3V9.75a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v.75M1 2.252l.621-.621A3 3 0 0 1 3.742.75h7.009m-.753 6h-6v-1.5a1.5 1.5 0 0 1 1.5-1.5h3a1.5 1.5 0 0 1 1.5 1.5v1.5ZM7 3.752v-3M7 9.75v5.998M4 12.75h5.998"/>`),
		g.Group(children),
	)
}