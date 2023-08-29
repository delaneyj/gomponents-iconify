package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SouthWestRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19q-.425 0-.713-.288T5 18v-8q0-.425.288-.713T6 9q.425 0 .713.288T7 10v5.6L17.9 4.7q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L8.4 17H14q.425 0 .713.288T15 18q0 .425-.288.713T14 19H6Z"/>`),
		g.Group(children),
	)
}