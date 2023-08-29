package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmstripSplitSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 1.5a.5.5 0 0 0-1 0v13a.5.5 0 0 0 1 0v-13ZM1 5.5A2.5 2.5 0 0 1 3.5 3h3v1h-3A1.5 1.5 0 0 0 2 5.5v5A1.5 1.5 0 0 0 3.5 12h3v1h-3A2.5 2.5 0 0 1 1 10.5v-5ZM9.5 4V3h3A2.5 2.5 0 0 1 15 5.5v5a2.5 2.5 0 0 1-2.5 2.5h-3v-1h3a1.5 1.5 0 0 0 1.5-1.5v-5A1.5 1.5 0 0 0 12.5 4h-3ZM12 5.5a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0v-1Zm.5 3.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 1 0v-1a.5.5 0 0 0-.5-.5ZM3 5.5a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0v-1ZM3.5 9a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 1 0v-1a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}