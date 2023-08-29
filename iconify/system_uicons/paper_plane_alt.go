package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.5 2.465l-8 8.033m3 8.002l-3-8.002l-7-2.998l15-5z"/>`),
		g.Group(children),
	)
}