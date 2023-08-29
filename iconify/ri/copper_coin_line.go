package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopperCoinLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.005 22.003c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0-12.95l4.95 4.95l-4.95 4.95l-4.95-4.95l4.95-4.95Zm0 2.829l-2.121 2.12l2.12 2.122l2.122-2.121l-2.121-2.121Z"/>`),
		g.Group(children),
	)
}