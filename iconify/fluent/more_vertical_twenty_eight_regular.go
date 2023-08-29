package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 8.25a2.25 2.25 0 1 1 0-4.5a2.25 2.25 0 0 1 0 4.5Zm0 8a2.25 2.25 0 1 1 0-4.5a2.25 2.25 0 0 1 0 4.5ZM11.75 22a2.25 2.25 0 1 0 4.5 0a2.25 2.25 0 0 0-4.5 0Z"/>`),
		g.Group(children),
	)
}