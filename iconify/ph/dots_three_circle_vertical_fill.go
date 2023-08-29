package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsThreeCircleVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 128a104 104 0 1 0-104 104a104.13 104.13 0 0 0 104-104ZM116 84a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm0 44a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm0 44a12 12 0 1 1 12 12a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}