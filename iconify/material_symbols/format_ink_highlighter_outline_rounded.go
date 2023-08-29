package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatInkHighlighterOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 24q-.825 0-1.413-.588T2 22q0-.825.588-1.413T4 20h16q.825 0 1.413.588T22 22q0 .825-.588 1.413T20 24H4Zm10.6-12L12 9.4l-4 4l2.575 2.6l4.025-4Zm-1.175-4L16 10.575L20 6.6L17.4 4l-3.975 4Zm-2.1-.725l5.4 5.4L12 17.425q-.6.6-1.413.6t-1.412-.6l-.125.1q-.275.225-.6.35T7.775 18h-3.05q-.35 0-.487-.3t.112-.55l2.3-2.275q-.6-.6-.625-1.438T6.6 12l4.725-4.725Zm0 0L16 2.6q.6-.6 1.413-.6t1.412.6l2.6 2.575q.6.6.6 1.413T21.425 8l-4.7 4.675l-5.4-5.4Z"/>`),
		g.Group(children),
	)
}