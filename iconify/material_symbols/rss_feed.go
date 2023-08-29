package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssFeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19q0-.825.588-1.413T5 17q.825 0 1.413.588T7 19q0 .825-.588 1.413T5 21Zm12 0q0-2.9-1.1-5.438t-3.013-4.45Q10.976 9.2 8.438 8.1T3 7V4q3.525 0 6.625 1.338t5.4 3.637q2.3 2.3 3.638 5.4T20 21h-3Zm-6 0q0-3.325-2.337-5.663T3 13v-3q2.3 0 4.3.863t3.488 2.35q1.487 1.487 2.35 3.487T14 21h-3Z"/>`),
		g.Group(children),
	)
}