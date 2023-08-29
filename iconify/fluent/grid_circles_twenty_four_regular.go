package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridCirclesTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 7a4 4 0 1 0-8 0a4 4 0 0 0 8 0ZM9.5 7a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM21 7a4 4 0 1 0-8 0a4 4 0 0 0 8 0Zm-1.5 0a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM7 21a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-1.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM21 17a4 4 0 1 0-8 0a4 4 0 0 0 8 0Zm-1.5 0a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/>`),
		g.Group(children),
	)
}