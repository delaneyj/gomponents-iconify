package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreHorizontalTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.75 12a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0Zm6 0a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0ZM18 13.75a1.75 1.75 0 1 0 0-3.5a1.75 1.75 0 0 0 0 3.5Z"/>`),
		g.Group(children),
	)
}