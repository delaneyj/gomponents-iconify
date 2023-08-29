package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="13.33" cy="29.75" r="2.25" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><circle cx="27" cy="29.75" r="2.25" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M22.57 7a7.52 7.52 0 0 1-.07-1a7.52 7.52 0 0 1 .07-1H11.49l.65 2Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M30 13.5h-.42L28.33 19h-15L8.76 4.53a1 1 0 0 0-.66-.65L4 2.62a1 1 0 1 0-.59 1.92L7 5.64l4.59 14.5l-1.64 1.34l-.13.13A2.66 2.66 0 0 0 9.74 25A2.75 2.75 0 0 0 12 26h16.69a1 1 0 0 0 0-2H11.84a.67.67 0 0 1-.56-1l2.41-2h15.44a1 1 0 0 0 1-.78l1.57-6.91a7.51 7.51 0 0 1-1.7.19Z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-5--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}