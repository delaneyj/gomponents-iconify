package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevereCold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 10q-.425 0-.713-.288T19 9q0-.425.288-.713T20 8q.425 0 .713.288T21 9q0 .425-.288.713T20 10ZM10 22v-3.6L7.4 21L6 19.6l4-4V14H8.4l-4 4L3 16.6L5.6 14H2v-2h3.6L3 9.4L4.4 8l4 4H10v-1.6l-4-4L7.4 5L10 7.6V4h2v3.6L14.6 5L16 6.4l-4 4V12h8v2h-3.6l2.6 2.6l-1.4 1.4l-4-4H12v1.6l4 4l-1.4 1.4l-2.6-2.6V22h-2Zm9-15V2h2v5h-2Z"/>`),
		g.Group(children),
	)
}