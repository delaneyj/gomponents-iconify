package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 5.25a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Zm0 4a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5ZM6.75 12a1.25 1.25 0 1 0 2.5 0a1.25 1.25 0 0 0-2.5 0Z"/>`),
		g.Group(children),
	)
}