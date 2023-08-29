package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuxIcons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="10" height="23.5" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5" ry="5"/><circle cx="24" cy="10.5" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="23.974" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.5" cy="10.5" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.5" cy="23.974" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="37.5" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="10.5" cy="37.5" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.5" cy="37.5" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}