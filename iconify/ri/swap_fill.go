package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.005 22.003c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Zm-5-13h2v4h2v-4h2l-3-3.5l-3 3.5Zm10 6h-2v-4h-2v4h-2l3 3.5l3-3.5Z"/>`),
		g.Group(children),
	)
}