package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UltimateSoftware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.1 16.2h7.8V24h-7.8zm0 15.6h7.8v7.8h-7.8zm7.8-23.4h7.8v7.8h-7.8zm7.8 7.8h7.8V24h-7.8zM12.3 8.4h7.8v7.8h-7.8zM27.9 24h7.8v7.8h-7.8zm-15.6 0h7.8v7.8h-7.8zm-7.8-7.8h7.8V24H4.5z"/>`),
		g.Group(children),
	)
}