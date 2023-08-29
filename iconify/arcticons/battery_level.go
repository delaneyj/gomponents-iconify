package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryLevel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.8 4.5h4.1c.7 0 1.1.4 1.1 1.1v1.9h6.8c1.2 0 2.2 1 2.2 2.2v30.6c0 1.2-1 2.2-2.2 2.2H14.2c-1.2 0-2.2-1-2.2-2.2V9.7c0-1.2 1-2.2 2.2-2.2h6.4V5.6c.1-.7.6-1.1 1.2-1.1Zm-2.8 3h9.7"/>`),
		g.Group(children),
	)
}