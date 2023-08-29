package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamStandOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7.925Q8 8.6 7.362 9.663T6.6 12L6 19h12l-.6-7q-.1-1.3-.738-2.363T15 7.9V10q0 1.25-.875 2.125T12 13q-1.25 0-2.125-.875T9 10V7.925ZM12 11q.425 0 .713-.288T13 10V6q0-.425-.288-.713T12 5q-.425 0-.713.288T11 6v4q0 .425.288.713T12 11Zm0-8q1.15 0 1.988.738t.987 1.862q1.85.8 3.063 2.438T19.4 11.8l.775 9.2H3.825l.775-9.2q.175-2.125 1.375-3.75t3.05-2.425q.15-1.125.988-1.875T12 3Z"/>`),
		g.Group(children),
	)
}