package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm0 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-1 3a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`),
		g.Group(children),
	)
}