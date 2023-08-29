package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pizzahut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 29.95l-10.02-6.91l-1.77-4.99H16.29l-1.77 4.99L4.5 29.95h39z"/>`),
		g.Group(children),
	)
}