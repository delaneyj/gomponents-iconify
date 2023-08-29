package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Via(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 31.73L25.37 7.75l1.65 8.96l-14.79 18.2l-7.73-3.18z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.97 22l.34-8.54l-6.14 1.08L4.5 31.73M25.37 7.75l6 32.5l12.13-2.67l-9.61-29.32l-8.52-.51z"/>`),
		g.Group(children),
	)
}