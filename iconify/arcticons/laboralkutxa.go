package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laboralkutxa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a10.61 10.61 0 1 0 10.61 10.61"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.51 43.5H24V25.72m0 12.99h3.96"/>`),
		g.Group(children),
	)
}