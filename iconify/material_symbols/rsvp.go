package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rsvp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 15V9h3.5q.625 0 1.063.438T6 10.5v1q0 .45-.263.825T5.1 12.9L6 15H4.5l-.85-2H2.5v2H1Zm6 0v-1.5h3v-.75H8q-.425 0-.713-.288T7 11.75V10q0-.425.288-.713T8 9h3.5v1.5h-3v.75h2q.425 0 .713.288t.287.712V14q0 .425-.288.713T10.5 15H7Zm7.25 0L12.5 9H14l1 3.425L16 9h1.5l-1.75 6h-1.5ZM18 15V9h3.5q.625 0 1.063.438T23 10.5v1q0 .625-.438 1.063T21.5 13h-2v2H18ZM2.5 11.5h2v-1h-2v1Zm17 0h2v-1h-2v1Z"/>`),
		g.Group(children),
	)
}