package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 12l3 3l3-3l-3-3zm15.5 2.5L21 12l-3-3l-2.5 2.5m-3.001-2.999L15 6l-3-3l-2.5 2.5M9 18l3 3l3-3l-3-3zM3 3l18 18"/>`),
		g.Group(children),
	)
}