package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V7a.5.5 0 0 1 1 0v1a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h1a.5.5 0 0 1 0 1H4Zm3 0a.5.5 0 0 1 0-1h2.5a.5.5 0 0 1 .5.5V5a.5.5 0 0 1-1 0V3.707L7.354 5.354a.5.5 0 1 1-.708-.708L8.293 3H7Z"/>`),
		g.Group(children),
	)
}