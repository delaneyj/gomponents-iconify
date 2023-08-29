package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddSubtractCircleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.5 12a.5.5 0 1 0 0 1h3a.5.5 0 0 0 0-1h-3ZM10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0-1a6.973 6.973 0 0 1-4.584-1.71l9.875-9.874A7 7 0 0 1 10 17ZM5.5 7a.5.5 0 0 1 .5-.5h1v-1a.5.5 0 0 1 1 0v1h1a.5.5 0 1 1 0 1H8v1a.5.5 0 0 1-1 0v-1H6a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}