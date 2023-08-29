package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolPercent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="18.439" cy="17.048" r="4.171"/><circle cx="29.561" cy="30.952" r="4.171"/><path d="m16.766 35.123l14.46-22.246"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}