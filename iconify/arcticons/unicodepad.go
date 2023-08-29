package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unicodepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.485 33.181a2.058 2.058 0 0 1-2.075-2.513l1.39-20.61c.116-1.715 1.423-3.449 4.518-3.404a3.692 3.692 0 0 1 3.87 4.19l-.36 5.524"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.316 40.677H6.755l1.907-13.888c2.76-20.107 34.085-14.994 32.749 5.185c-8.383 4.177-28.296.523-32.75-5.185"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.316 38.66v2.017h10.083V38.66s.203-4.477-5.042-4.438s-5.041 4.438-5.041 4.438M20.762 14.121l.24-6.005m16.342 11.422l4.331-3.834"/><circle cx="16.166" cy="21.85" r="1.467" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.61" cy="25.884" r="1.467" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.435 31.392l-.116 2.733"/>`),
		g.Group(children),
	)
}