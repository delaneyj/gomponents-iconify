package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProcessLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 40h10M4 32h20m-2 8h5m10 0h7m-9-8h9m-12-8h12m-28 0h6M4 24h2m-2-8h4M4 8h7l8 8h25M22 8h22"/>`),
		g.Group(children),
	)
}