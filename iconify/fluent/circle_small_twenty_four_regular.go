package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleSmallTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 9.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM8 12a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z"/>`),
		g.Group(children),
	)
}