package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Follow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="38.551" r="3.644" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="9.144" r="3.644" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14.851" cy="23.848" r="3.644" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.149" cy="23.848" r="3.644" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="7.265" cy="40.275" r="2.225" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="40.735" cy="40.275" r="2.225" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}