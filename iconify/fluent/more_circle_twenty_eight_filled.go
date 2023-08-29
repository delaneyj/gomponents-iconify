package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreCircleTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2C7.373 2 2 7.373 2 14s5.373 12 12 12s12-5.373 12-12S20.627 2 14 2ZM9.5 14a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm6 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm6 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}