package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.754 4a2.75 2.75 0 0 0-2.75 2.75V9H24.5v14.25a.75.75 0 0 0 1.5 0V6.75A2.75 2.75 0 0 0 23.25 4H4.755ZM14 10.5H2.004v9.75A3.75 3.75 0 0 0 5.754 24h4.496A3.75 3.75 0 0 0 14 20.25V10.5Zm-8.5 4.25a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}