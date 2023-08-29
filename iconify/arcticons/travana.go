package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Travana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.617" cy="35.794" r="6.307" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="22.535" cy="21.662" r="4.768" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.751" cy="18.569" r="4.768" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.869" cy="39.113" r="4.387" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.619" cy="28.552" r="3.309" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.116" cy="29.334" r="3.309" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.615" cy="7.511" r="3.011" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="25.905" cy="10.94" r="3.878" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}