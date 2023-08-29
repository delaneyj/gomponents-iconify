package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 6.5A1.75 1.75 0 1 1 10 3a1.75 1.75 0 0 1 0 3.5ZM10 17a1.75 1.75 0 1 1 0-3.5a1.75 1.75 0 0 1 0 3.5Zm-1.75-7a1.75 1.75 0 1 0 3.5 0a1.75 1.75 0 0 0-3.5 0Z"/>`),
		g.Group(children),
	)
}