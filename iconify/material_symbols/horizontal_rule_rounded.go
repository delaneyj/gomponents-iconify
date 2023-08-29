package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalRuleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 13q-.425 0-.713-.288T4 12q0-.425.288-.713T5 11h14q.425 0 .713.288T20 12q0 .425-.288.713T19 13H5Z"/>`),
		g.Group(children),
	)
}