package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackpackOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V8q0-1.4.85-2.45T7 4.15V2h3v2h4V2h3v2.15q1.3.35 2.15 1.4T20 8v14H4Zm2-2h12V8q0-.825-.588-1.413T16 6H8q-.825 0-1.413.588T6 8v12Zm8.5-4h2v-4h-9v2h7v2ZM12 13Z"/>`),
		g.Group(children),
	)
}