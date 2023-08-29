package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotchesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 40v152a8 8 0 0 1-8 8H40a8 8 0 0 1-5.66-13.66l152-152A8 8 0 0 1 200 40Z"/>`),
		g.Group(children),
	)
}