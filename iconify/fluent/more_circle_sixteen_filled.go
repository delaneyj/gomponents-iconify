package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreCircleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12Zm0-5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM5 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}