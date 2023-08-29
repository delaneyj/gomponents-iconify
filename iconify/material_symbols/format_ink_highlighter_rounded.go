package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatInkHighlighterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 24q-.825 0-1.413-.588T2 22q0-.825.588-1.413T4 20h16q.825 0 1.413.588T22 22q0 .825-.588 1.413T20 24H4Zm6.6-16l5.4 5.425l-4 4q-.6.6-1.413.6t-1.412-.6l-.125.1q-.275.225-.6.35T7.775 18h-3.05q-.35 0-.487-.3t.112-.55l2.3-2.275q-.6-.6-.625-1.438T6.6 12l4-4ZM12 6.575L16 2.6q.6-.6 1.413-.6t1.412.6l2.6 2.575q.6.6.6 1.413T21.425 8l-4 4L12 6.575Z"/>`),
		g.Group(children),
	)
}