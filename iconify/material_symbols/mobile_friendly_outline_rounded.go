package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileFriendlyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 23q-.825 0-1.413-.588T4 21V3q0-.825.588-1.413T6 1h10q.825 0 1.413.588T18 3v4h-2V6H6v12h10v-1h2v4q0 .825-.588 1.413T16 23H6Zm0-3v1h10v-1H6Zm8.95-6.8l4.95-4.95q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-5.65 5.65q-.3.3-.7.3t-.7-.3l-2.85-2.85q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l2.15 2.15ZM6 4h10V3H6v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}