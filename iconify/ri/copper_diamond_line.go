package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopperDiamondLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.005 22.003c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm-3-12h6l2.5 3.5l-5.5 5.5l-5.5-5.5l2.5-3.5Zm1.03 2l-.92 1.29l2.89 2.89l2.89-2.89l-.92-1.29h-3.94Z"/>`),
		g.Group(children),
	)
}