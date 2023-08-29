package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.854 2.15a.5.5 0 0 1 0 .706l-15 15a.5.5 0 0 1-.708-.707l15-15a.5.5 0 0 1 .708 0Z"/>`),
		g.Group(children),
	)
}