package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunSnow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9a3 3 0 1 0 0 6m-8-3h1m11 9V3m-4 1V3m0 18v-1m-6.36-1.64l.7-.7m0-11.32l-.7-.7M14 12h8m-5-8l-3 3m0 10l3 3m4-5l-3-3l3-3"/>`),
		g.Group(children),
	)
}