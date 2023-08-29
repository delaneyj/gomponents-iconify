package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 2a4 4 0 1 0 0 8a4 4 0 0 0 0-8ZM1 6a5 5 0 1 1 10 0A5 5 0 0 1 1 6Z"/>`),
		g.Group(children),
	)
}