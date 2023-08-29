package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatInkHighlighter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 24v-4h20v4H2Zm8.6-16l5.4 5.425l-4 4q-.6.6-1.413.6t-1.412-.6L8.5 18h-5l3.15-3.125q-.6-.6-.625-1.438T6.6 12l4-4ZM12 6.575L16 2.6q.6-.6 1.413-.6t1.412.6l2.6 2.575q.6.6.6 1.413T21.425 8l-4 4L12 6.575Z"/>`),
		g.Group(children),
	)
}