package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.784 2.22a.75.75 0 0 1 0 1.06L3.28 21.784a.75.75 0 1 1-1.06-1.06L20.723 2.22a.75.75 0 0 1 1.06 0Z"/>`),
		g.Group(children),
	)
}