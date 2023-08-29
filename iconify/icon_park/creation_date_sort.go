package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreationDateSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 5V30.0036H42V5"/><path stroke-linejoin="round" d="M30 37L24 43L18 37"/><path stroke-linejoin="round" d="M24 30V43"/><path d="M16.0009 15.001L32.0009 15.001"/><path d="M24.0005 7V23"/></g>`),
		g.Group(children),
	)
}