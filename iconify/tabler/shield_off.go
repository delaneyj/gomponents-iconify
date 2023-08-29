package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.67 17.667A12 12 0 0 1 12 21A12 12 0 0 1 3.5 6c.794.036 1.583-.006 2.357-.124m3.128-.926A11.997 11.997 0 0 0 12 3a12 12 0 0 0 8.5 3a12 12 0 0 1-1.116 9.376M3 3l18 18"/>`),
		g.Group(children),
	)
}