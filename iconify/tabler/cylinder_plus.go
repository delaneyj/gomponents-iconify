package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CylinderPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 6a7 3 0 1 0 14 0A7 3 0 1 0 5 6"/><path d="M5 6v12c0 1.657 3.134 3 7 3c.173 0 .345-.003.515-.008M19 12V6m-3 13h6m-3-3v6"/></g>`),
		g.Group(children),
	)
}