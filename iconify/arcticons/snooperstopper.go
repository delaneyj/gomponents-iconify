package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snooperstopper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.68 12.92a21.19 21.19 0 0 1 13.81 10c-9.42 16.33-29.63 16.2-39 0a21.35 21.35 0 0 1 13.82-9.95A10.08 10.08 0 0 0 24 31.33h0a10.08 10.08 0 0 0 10.07-10.08h0a10.1 10.1 0 0 0-4.4-8.33Zm-4.32 8.4c0 3.61 3.19.77 3.42.55c.5-.5 1.6 0 1 1c-3.22 5.57-3.92 5-8 5c-2.3 0-2.3-2.64-2.3-6.58m0-3.75v3.75m1.89-5.85v5.85m2-7v7m2.07-5.84v5.84"/>`),
		g.Group(children),
	)
}