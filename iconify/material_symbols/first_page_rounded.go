package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstPageRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18q-.425 0-.713-.288T6 17V7q0-.425.288-.713T7 6q.425 0 .713.288T8 7v10q0 .425-.288.713T7 18Zm6.8-6l3.9 3.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-4.6-4.6q-.15-.15-.212-.325T11.425 12q0-.2.063-.375t.212-.325l4.6-4.6q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L13.8 12Z"/>`),
		g.Group(children),
	)
}