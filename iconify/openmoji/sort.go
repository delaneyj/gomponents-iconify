package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M16 33h31v6H16zm0-12h39.986v6H16zm0 24h26v6H16z"/><path fill="#D0CFCE" d="M16 24h40v3H16zm0 12h31v3H16zm.333 12h26v3h-26z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M16 21h40v6H16zm0 12h31v6H16zm.333 12h26v6h-26z"/>`),
		g.Group(children),
	)
}