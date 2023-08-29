package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11.933 5H5v16h13v-8m-4 4H9"/><path d="M9 13h5V9H9zm6-8V3m3 3l2-2m-1 5h2"/></g>`),
		g.Group(children),
	)
}