package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CdSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 6a2 2 0 1 1 .001 3.999A2 2 0 0 1 8 6Zm0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm0-7a6 6 0 1 0 0 12A6 6 0 0 0 8 2Z"/>`),
		g.Group(children),
	)
}