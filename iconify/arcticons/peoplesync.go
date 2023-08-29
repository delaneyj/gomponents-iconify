package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peoplesync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39 10.947S53.018 27.61 21.146 41.78"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.52 32.67l-8.915 9.83l16.63-1.65m-23.001 1.09c-7.803.284-16.06-25.212 13.783-32.88"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.633 5.5l11.82 3.167l-8.317 8.317"/>`),
		g.Group(children),
	)
}