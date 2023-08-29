package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g class="icon-tabler" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 18V6a2 2 0 0 0-2.75-1.84L9 10.26a2 2 0 0 0 0 3.5l8.25 6.1A2 2 0 0 0 20 18.11"/><path d="M4 20V4"/></g>`),
		g.Group(children),
	)
}