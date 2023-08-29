package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatInkHighlighterOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 24v-4h20v4H2Zm12.6-12L12 9.4l-4 4l2.575 2.6l4.025-4Zm-1.175-4L16 10.575L20 6.6L17.4 4l-3.975 4Zm-2.1-.725l5.4 5.4L12 17.425q-.6.6-1.413.6t-1.412-.6L8.5 18h-5l3.15-3.125q-.6-.6-.625-1.438T6.6 12l4.725-4.725Zm0 0L16 2.6q.6-.6 1.413-.6t1.412.6l2.6 2.575q.6.6.6 1.413T21.425 8l-4.7 4.675l-5.4-5.4Z"/>`),
		g.Group(children),
	)
}