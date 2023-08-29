package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PigOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11v.01M10 6h1.499l4.5-3v3.803A6.019 6.019 0 0 1 18.657 10h1.341a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-1.342c-.057.16-.12.318-.19.472M16.999 17v1.5a1.5 1.5 0 0 1-3 0v-.583a6.04 6.04 0 0 1-1 .083h-4a6.04 6.04 0 0 1-1-.083v.583a1.5 1.5 0 0 1-3 0v-2.027a6 6 0 0 1 1.5-9.928M3 3l18 18"/>`),
		g.Group(children),
	)
}