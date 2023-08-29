package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyZone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 39.284h6.324m4.217 0H43.5m-32.676 0L17.15 8.716h13.7m1.022 4.94l5.303 25.627M24.527 13.657v15.355"/><circle cx="24.527" cy="32.965" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}