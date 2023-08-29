package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkEraseOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 13.4l-1.9 1.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l1.9-1.9l-1.9-1.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.9 1.9l1.9-1.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L18.4 12l1.9 1.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L17 13.4ZM6 23q-.825 0-1.413-.588T4 21V3q0-.825.588-1.413T6 1h10q.825 0 1.413.588T18 3v4h-2V6H6v12h10v-1h2v4q0 .825-.588 1.413T16 23H6Zm0-3v1h10v-1H6ZM6 4h10V3H6v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}