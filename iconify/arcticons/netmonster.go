package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Netmonster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.273" cy="19.535" r="2.376" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.186 34.085A15.677 15.677 0 0 1 19.526 4.5m9.435.152a15.678 15.678 0 0 1 1.286 29.251"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.597 29.266a9.502 9.502 0 0 1 1.418-18.134m4.481.098a9.502 9.502 0 0 1 1.375 17.883"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.273 21.91L14.77 40.916a18.756 18.756 0 0 0 19.004 0Z"/>`),
		g.Group(children),
	)
}