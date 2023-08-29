package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesPants(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M33 44H42L38 4H10L6 44H15L24 19L33 44Z"/><path d="M24 4V9.5"/><path d="M17.0004 4C17.0004 4 17.0004 10 15.0004 12C13.0004 14 8.90039 15 8.90039 15"/><path d="M31 4C31 4 31 10 33 12C35 14 39.1 15 39.1 15"/></g>`),
		g.Group(children),
	)
}