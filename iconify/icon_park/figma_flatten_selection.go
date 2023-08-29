package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaFlattenSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 18L24 26L16 18"/><path d="M36 42H12"/><path d="M42 34H6"/><path d="M24 6V26"/></g>`),
		g.Group(children),
	)
}