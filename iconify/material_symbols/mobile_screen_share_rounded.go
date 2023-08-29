package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileScreenShareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 15q-.425 0-.713-.288T9 14v-.5q0-1.25.875-2.125T12 10.5h1V9.1q0-.175.15-.238t.275.063l2.225 2.225q.15.15.15.35t-.15.35l-2.225 2.225q-.125.125-.275.062T13 13.9v-1.4h-1q-.425 0-.713.287T11 13.5v.5q0 .425-.288.713T10 15Zm-3 8q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}