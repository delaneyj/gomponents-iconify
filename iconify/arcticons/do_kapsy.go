package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoKapsy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.486 26.958V43.5H8.261v-39h23.225v2.666"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.646 14.798c6.025 3.213 14.497 3.213 22.093-4.054M17.208 26.958c2.775-3.47 10.663-11.357 22.239-4.82"/><circle cx="26.775" cy="11.803" r="3.287" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}