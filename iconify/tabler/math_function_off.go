package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathFunctionOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10h1c.882 0 .986.777 1.694 2.692M13 17c.864 0 1.727-.663 2.495-1.512m1.717-2.302C18.205 11.736 19.602 10 21 10M3 19c0 1.5.5 2 2 2s2-4 3-9c.237-1.186.446-2.317.647-3.35m.727-3.248C9.797 3.91 10.284 3 11 3c1.5 0 2 .5 2 2m-8 7h6M3 3l18 18"/>`),
		g.Group(children),
	)
}