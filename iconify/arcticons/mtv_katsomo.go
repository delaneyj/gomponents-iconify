package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MtvKatsomo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.45 18.913a14.455 14.455 0 1 1-.002-.116M34.145 37.31l-2.357 6.123l-2.356-6.123m-15.577 2.451a2.332 2.332 0 1 1 4.664 0V43.5m-4.664-6.076V43.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.52 39.761a2.332 2.332 0 1 1 4.663 0V43.5m2.77-7.81v6.669a1.05 1.05 0 0 0 1.112 1.111h.333m-2.556-5.891h2.334M20.84 25.493a21.634 21.634 0 0 0 9.213-6a21.633 21.633 0 0 0-9.212-6h0a20.5 20.5 0 0 0 0 12Z"/>`),
		g.Group(children),
	)
}