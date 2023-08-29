package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inviziblepro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.21" rx="19.5" ry="12.01"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.21" d="M40 40L8 8"/><circle cx="24" cy="24" r="8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.21"/><circle cx="24" cy="24" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.21"/>`),
		g.Group(children),
	)
}