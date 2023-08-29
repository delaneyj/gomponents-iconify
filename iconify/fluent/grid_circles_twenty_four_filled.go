package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridCirclesTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm10 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm-6 6a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm6 4a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`),
		g.Group(children),
	)
}