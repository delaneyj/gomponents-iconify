package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.15 7.11L4.5 30.75l5.85 10.14h27.3l5.85-10.14L29.85 7.11ZM4.5 30.75h27.29m5.86 10.14L24 17.25m5.85-10.14L16.21 30.75"/>`),
		g.Group(children),
	)
}