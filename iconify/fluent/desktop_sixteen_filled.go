package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.501 2a1.5 1.5 0 0 0-1.5 1.5v6.997a1.5 1.5 0 0 0 1.5 1.5h2.5V13H4.495a.5.5 0 0 0 0 1H11.5a.5.5 0 1 0 0-1H10v-1.003h2.501a1.5 1.5 0 0 0 1.5-1.5V3.5a1.5 1.5 0 0 0-1.5-1.5h-9Zm5.5 9.997V13H7v-1.003h2Z"/>`),
		g.Group(children),
	)
}