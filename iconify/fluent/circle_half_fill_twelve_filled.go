package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfFillTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 6a5 5 0 1 1 10 0A5 5 0 0 1 1 6Zm1.5 0h7a3.5 3.5 0 1 0-7 0Z"/>`),
		g.Group(children),
	)
}