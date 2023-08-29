package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandWish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 6l5.981 2.392l-.639 6.037c-.18.893.06 1.819.65 2.514A3 3 0 0 0 10.373 18a4.328 4.328 0 0 0 4.132-3.57c-.18.893.06 1.819.65 2.514A3 3 0 0 0 17.535 18a4.328 4.328 0 0 0 4.132-3.57L22 9.797m-7.496 4.632l.334-3"/>`),
		g.Group(children),
	)
}