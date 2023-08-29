package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prepaidguthaben(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="29.88" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="13.62" ry="13.633"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.12 37.633A13.626 13.626 0 0 1 4.5 24h0a13.626 13.626 0 0 1 13.62-13.633"/>`),
		g.Group(children),
	)
}