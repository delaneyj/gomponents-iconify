package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLineSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 8A6 6 0 1 1 2 8a6 6 0 0 1 12 0Zm-1.025.5h-9.95a5 5 0 0 0 9.95 0Zm0-1a5 5 0 0 0-9.95 0h9.95Z"/>`),
		g.Group(children),
	)
}