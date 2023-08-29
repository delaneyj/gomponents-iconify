package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<rect width="192" height="192" x="32" y="32" fill="currentColor" rx="16"/>`),
		g.Group(children),
	)
}