package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneHundredFourUa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.5 12.75l3.671-2v14.687M24.95 12.48l-7.44 11.227"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M21.236 25.436a4.864 4.864 0 0 1-4.864-4.864v-4.957a4.864 4.864 0 1 1 9.729 0v4.957a4.864 4.864 0 0 1-4.865 4.864Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.653 25.436V10.75l-7.883 9.866h9.73"/><circle cx="17.212" cy="36.968" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.941 28.436v5.467a3.313 3.313 0 1 0 6.626.001v-5.468m.001 5.467v3.313m10.091-3.313a3.312 3.312 0 0 1-3.311 3.313h-.002a3.312 3.312 0 0 1-3.313-3.311v-2.156a3.312 3.312 0 0 1 3.312-3.313h.001a3.312 3.312 0 0 1 3.313 3.312v.001m.001 5.467v-8.78"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}