package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.51 19.5l13 7.5l-13 7.5v-15Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-21a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v21a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Zm-6.61-26.13V9.62a1.12 1.12 0 0 0-1.12-1.12H12.23a1.12 1.12 0 0 0-1.12 1.12v1.75m28.78 3.13v-2a1.12 1.12 0 0 0-1.12-1.12H9.23a1.12 1.12 0 0 0-1.12 1.12v2"/>`),
		g.Group(children),
	)
}