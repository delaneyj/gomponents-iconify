package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserDeleteOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 14h-6a1 1 0 1 1 0-2h6a1 1 0 1 1 0 2zM9 6c1.654 0 3 1.346 3 3s-1.346 3-3 3s-3-1.346-3-3s1.346-3 3-3m0-2C6.236 4 4 6.238 4 9s2.236 5 5 5s5-2.238 5-5s-2.236-5-5-5zm0 13c2.021 0 3.301.771 3.783 1.445C12.1 18.705 10.814 19 9 19c-1.984 0-3.206-.305-3.818-.542C5.641 17.743 6.959 17 9 17m0-2c-3.75 0-6 2-6 4c0 1 2.25 2 6 2c3.518 0 6-1 6-2c0-2-2.354-4-6-4z"/>`),
		g.Group(children),
	)
}