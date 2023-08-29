package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarPatternSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 2A2.5 2.5 0 0 0 2 4.5v7A2.5 2.5 0 0 0 4.5 14h7a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 11.5 2h-7ZM3 4.5A1.5 1.5 0 0 1 4.5 3h7A1.5 1.5 0 0 1 13 4.5V6H3V4.5ZM11.707 7H13v.293L10.293 10H8.707l3-3Zm-1.414 0l-3 3H5.707l3-3h1.586Zm-6 3H3v-.293L5.707 7h1.586l-3 3ZM3 8.293V7h1.293L3 8.293ZM11.707 10L13 8.707V10h-1.293Z"/>`),
		g.Group(children),
	)
}