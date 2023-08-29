package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediumAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="13.979" cy="24" r="10.479" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="32.318" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.934" ry="10.479"/><ellipse cx="42.339" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.161" ry="10.479"/>`),
		g.Group(children),
	)
}