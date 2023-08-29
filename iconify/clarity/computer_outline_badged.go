package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M1 25v3.4A2.6 2.6 0 0 0 3.6 31h28.74a2.6 2.6 0 0 0 2.6-2.6V25Zm32 3.4a.6.6 0 0 1-.6.6H3.56a.6.6 0 0 1-.6-.6v-1.87h9.95a1.64 1.64 0 0 0 1.5 1h7.13a1.64 1.64 0 0 0 1.5-1H33Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M22.5 6a7.52 7.52 0 0 1 .07-1H5.5A1.5 1.5 0 0 0 4 6.5V23h2V7h16.57a7.52 7.52 0 0 1-.07-1Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M30 13.5V23h2v-9.78a7.49 7.49 0 0 1-2 .28Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><path fill="currentColor" d="M23.13 9H8v13.88h1.6V10.6h14.48a7.49 7.49 0 0 1-.95-1.6Z" class="clr-i-outline--badged clr-i-outline-path-4--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-5--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}