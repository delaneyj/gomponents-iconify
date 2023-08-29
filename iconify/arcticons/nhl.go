package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nhl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5c-2.272 1.893-6.579 4.45-11.36 1.136a10.131 10.131 0 0 1-4.733 4.118c1.704 3.644 1.468 4.449 1.468 7.241s-2.603 8-2.603 11.502c0 3.944 2.035 11.335 12.306 12.211c2.013.172 3.62 1.732 4.922 2.792m0-39c2.272 1.893 6.579 4.45 11.36 1.136a10.131 10.131 0 0 0 4.733 4.118c-1.704 3.644-1.468 4.449-1.468 7.241s2.603 8 2.603 11.502c0 3.944-2.035 11.335-12.306 12.211c-2.013.172-3.62 1.732-4.922 2.792M36.086 6.637L6.828 29.65m32.224-10.096L13.767 39.443"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.914 14.255v8.291l4.145-3.342M14.338 35.465v-8.291l5.493 3.863v-8.291m3.158-2.388v8.291m4.797-12.174v8.29m-4.797-.277l4.797-3.883m-4.797-.247l-1.05.847m8.975-6.95l-1.05.847M14.338 27.174l-1.05.847"/>`),
		g.Group(children),
	)
}