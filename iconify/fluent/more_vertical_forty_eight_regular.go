package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24.002 15.75a3.25 3.25 0 1 1 0-6.5a3.25 3.25 0 0 1 0 6.5Zm0 11.5a3.25 3.25 0 1 1 0-6.5a3.25 3.25 0 0 1 0 6.5Zm-3.25 8.25a3.25 3.25 0 1 0 6.5 0a3.25 3.25 0 0 0-6.5 0Z"/>`),
		g.Group(children),
	)
}