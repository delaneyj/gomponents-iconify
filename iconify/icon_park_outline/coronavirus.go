package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coronavirus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7.635 35.5A20.11 20.11 0 0 0 13 40.706m14.868 2.92A20.107 20.107 0 0 1 24 44c-1.324 0-2.617-.129-3.869-.374M43.55 28.243a20.173 20.173 0 0 0 .017-8.4M40.365 12.5A20.102 20.102 0 0 0 35 7.294M20.155 4.37A20.1 20.1 0 0 1 24 4c1.315 0 2.6.127 3.845.37M7.635 12.5A20.104 20.104 0 0 1 13 7.294m27.5 28a20.096 20.096 0 0 1-5.365 5.206M16 24H4m6 14l8.343-8.343M24 32v12m14-6l-8.343-8.343M32 24h12m-6-14l-8.343 8.343M24 16V4m-14 6l8.343 8.343m-13.91 1.5A20.083 20.083 0 0 0 4 24c0 1.425.15 2.816.433 4.157M24 32a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/>`),
		g.Group(children),
	)
}