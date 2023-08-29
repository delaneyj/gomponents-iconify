package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Services(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="9.379" cy="9.379" r="3.879" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.379" cy="38.621" r="3.879" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.379" cy="24" r="3.879" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="25.064" height="7.758" x="17.436" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.79" ry="1.79"/><rect width="25.064" height="7.758" x="17.436" y="34.742" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.79" ry="1.79"/><rect width="25.064" height="7.758" x="17.436" y="20.121" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.79" ry="1.79"/>`),
		g.Group(children),
	)
}