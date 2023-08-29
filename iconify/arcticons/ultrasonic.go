package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ultrasonic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.504 28.487a3.817 3.817 0 1 1 7.633 0m-.001 7.268a3.817 3.817 0 0 1-7.632 0m0 0v-7.268m7.632 0v7.268m-7.632-7.268v-5.213m-.006.005c0-9.376 8.283-16.977 18.5-16.977s18.5 7.6 18.5 16.977m0 5.213a3.817 3.817 0 1 0-7.633 0m.001 7.268a3.817 3.817 0 1 0 7.633 0m-.001 0v-7.268m-7.632 0v7.268m7.632-7.268V23.28"/>`),
		g.Group(children),
	)
}