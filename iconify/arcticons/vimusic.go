package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vimusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.225 24.903a98.832 98.832 0 0 0-12.78-12.65q-2.701-2.038-.181-3.765l2.931-1.717a4.2 4.2 0 0 1 5.713.623a80.332 80.332 0 0 1 10.5 16.746l-.33-17.62a2.714 2.714 0 0 1 2.85-.923a22.97 22.97 0 0 0 11.385 3.584q1.707 0 1.637 1.837l-.13 6.676A1.334 1.334 0 0 1 40.313 19a25.321 25.321 0 0 1-10.802-2.801l.04 9.637v7.67a8.697 8.697 0 0 1-11.967 8.374c-8.373-3.293-6.626-15.893 2.64-16.977Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.145 33.617q1.957-.602 5.01 2.992c3.001 3.192-.954 5.27-3.103 2.7q-3.082-3.262-1.907-5.692Z"/>`),
		g.Group(children),
	)
}