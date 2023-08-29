package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProtonVpnAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="10.44" cy="17.38" r="4.38" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.56" cy="8.88" r="4.38" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.41" cy="39.12" r="4.38" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.19 20.78l12.06 15.31m-10.58-19.8l18.7-5.93m2.92 2.7l-6.45 21.93"/>`),
		g.Group(children),
	)
}