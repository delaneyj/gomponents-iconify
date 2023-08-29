package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeFsSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.25 9H13V7h1.25a.5.5 0 1 0 0-1H13V5a.5.5 0 0 0-1 0v1h-2V5a.5.5 0 0 0-1 0v1H8a.5.5 0 0 0 0 1h1v2H8a.5.5 0 0 0 0 1h1v1a.5.5 0 0 0 1 0v-1h2v1a.5.5 0 0 0 1 0v-1h1.25a.5.5 0 1 0 0-1ZM10 9V7h2v2h-2ZM6.5 4.5A.5.5 0 0 1 6 5H3v2.5h2.5a.5.5 0 0 1 0 1H3v3a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5H6a.5.5 0 0 1 .5.5Z"/>`),
		g.Group(children),
	)
}