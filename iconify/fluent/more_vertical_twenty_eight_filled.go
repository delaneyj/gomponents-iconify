package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 8.75a2.75 2.75 0 1 1 0-5.5a2.75 2.75 0 0 1 0 5.5Zm0 8a2.75 2.75 0 1 1 0-5.5a2.75 2.75 0 0 1 0 5.5ZM11.25 22a2.75 2.75 0 1 0 5.5 0a2.75 2.75 0 0 0-5.5 0Z"/>`),
		g.Group(children),
	)
}