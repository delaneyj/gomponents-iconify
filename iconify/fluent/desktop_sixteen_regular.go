package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v5.997a2 2 0 0 0 2 2h2.005v1.011H4.506a.5.5 0 0 0 0 1h6.996a.5.5 0 1 0 0-1h-1.5v-1.01H12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4Zm5.003 9.997v1.011H7.005v-1.01h1.998ZM3 4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v5.997a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4Z"/>`),
		g.Group(children),
	)
}