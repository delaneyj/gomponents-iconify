package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seabank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.097 32.244c-4.448-4.767-5.246-11.808-1.974-17.417c3.27-5.61 9.862-8.503 16.306-7.158c6.443 1.345 11.263 6.62 11.922 13.051c.66 6.43-2.992 12.542-9.033 15.12"/><path d="M9.403 32.784C3.717 24.96 5.345 14.122 13.092 8.241c7.746-5.88 18.842-4.706 25.133 2.66c6.29 7.367 5.528 18.294-1.728 24.75"/><path d="M30.346 35.837c-2.293 1-7.966-4.683-17.216-3.57m-3.35 8.794c.54-.547 4.97-4.822 14.726.4c7.028 3.326 12.375 1.886 14.32.554"/><path d="M5.284 39.077s6.294-8.148 20.055-1.695c7.284 4.047 13.558 1.33 17.377-1.411"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.308 26.974c.776 1.01 1.75 1.386 3.102 1.386h1.873a3.156 3.156 0 0 0 3.156-3.156v-.014a3.156 3.156 0 0 0-3.156-3.156h-2.066a3.16 3.16 0 0 1-3.159-3.159h0a3.166 3.166 0 0 1 3.166-3.166h1.863c1.353 0 2.326.376 3.102 1.386"/>`),
		g.Group(children),
	)
}