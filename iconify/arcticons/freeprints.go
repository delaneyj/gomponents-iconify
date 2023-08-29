package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freeprints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.503 9.39A9.484 9.484 0 0 1 24 18.162c-3.77 0-1.92 10.564 0 14.45"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.09 18.422s-4.113-7.086-11.538-8.786s-8.251 6.496-5.846 10.27s5.081 6.303 8.642 5.975"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.194 30.577s.698 3.955-2.012 6.345c-2.643 2.332-8.09 2.264-10.444-.309c-1.9-2.076-2.31-6.24.476-8.439s10.735-7.482 10.735-7.482M29.497 9.39A9.484 9.484 0 0 0 24 18.162c3.77 0 1.92 10.564 0 14.45"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.91 18.422s4.113-7.086 11.538-8.786s8.251 6.496 5.846 10.27s-5.081 6.303-8.642 5.975"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.806 30.577s-.698 3.955 2.012 6.345c2.643 2.332 8.09 2.264 10.444-.309c1.9-2.076 2.31-6.24-.476-8.439S26.05 20.692 26.05 20.692"/>`),
		g.Group(children),
	)
}