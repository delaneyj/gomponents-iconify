package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAssistant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="15.57" cy="17.82" r="10.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.36" cy="22.27" r="4.46" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.83" cy="16.6" r="2.67" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.36" cy="34.8" r="5.45" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}