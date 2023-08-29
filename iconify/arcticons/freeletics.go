package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freeletics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.801 29.306L38.496 9.504L18.694 4.199L4.199 18.694l5.305 19.802l19.802 5.305l14.495-14.495zM13 39.432V9.893"/>`),
		g.Group(children),
	)
}