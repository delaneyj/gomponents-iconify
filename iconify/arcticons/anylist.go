package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anylist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="21.33" height="33.044" x="13.335" y="7.478" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.532" ry="3.532"/><rect width="13.278" height="3.581" x="17.361" y="14.638" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.369" ry="3.532"/><rect width="13.278" height="3.581" x="17.361" y="21.421" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.369" ry="3.532"/><rect width="13.278" height="3.581" x="17.361" y="28.596" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.369" ry="3.532"/>`),
		g.Group(children),
	)
}