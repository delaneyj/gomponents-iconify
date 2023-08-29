package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextMining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 28h8v2h-8zm0-4h12v2H18zm0-4h12v2H18z"/><path fill="currentColor" d="M16.001 26.473H16L4.284 12.955L9.5 6h13l5.216 6.955l-3.24 3.737l1.513 1.31l4.295-4.957L23.5 4h-15l-6.784 9.045l12.772 14.737l1.513-1.309z"/>`),
		g.Group(children),
	)
}