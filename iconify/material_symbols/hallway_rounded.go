package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HallwayRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V6q0-.825.588-1.413T5 4h3L11.3.7q.3-.3.7-.3t.7.3L16 4h3q.825 0 1.413.588T21 6v14q0 .825-.588 1.413T19 22H5Zm6.25-5L9.4 14.525q-.15-.2-.4-.2t-.4.2l-2 2.675q-.2.25-.05.525T7 18h10q.3 0 .45-.275t-.05-.525l-2.75-3.675q-.15-.2-.4-.2t-.4.2L11.25 17ZM10.1 4h3.8L12 2.1L10.1 4Z"/>`),
		g.Group(children),
	)
}