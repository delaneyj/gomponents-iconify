package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BouyguesTelecom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="12.97" cy="29.193" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.786" ry="12.452" transform="rotate(-34.12 12.97 29.193)"/><ellipse cx="35.03" cy="29.485" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="12.452" ry="5.786" transform="rotate(-55.88 35.03 29.485)"/><ellipse cx="23.854" cy="13.492" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="12.452" ry="5.786" transform="rotate(-.263 23.846 13.496)"/>`),
		g.Group(children),
	)
}