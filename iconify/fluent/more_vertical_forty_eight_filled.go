package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 16a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0 11.5a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm-3.5 8a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Z"/>`),
		g.Group(children),
	)
}