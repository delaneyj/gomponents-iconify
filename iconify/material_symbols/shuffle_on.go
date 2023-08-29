package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23q-.825 0-1.413-.588T1 21V3q0-.825.588-1.413T3 1h18q.825 0 1.413.588T23 3v18q0 .825-.588 1.413T21 23H3Zm11-3h6v-6h-2v2.55l-3.175-3.175L13.4 14.8l3.2 3.2H14v2Zm-8.6 0L18 7.4V10h2V4h-6v2h2.6L4 18.6L5.4 20Zm3.775-9.425l1.4-1.4L5.4 4L4 5.4l5.175 5.175Z"/>`),
		g.Group(children),
	)
}