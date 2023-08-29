package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yaata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39 9.512H9a2 2 0 0 0-2 2v26.976l6.977-6.976H39a2 2 0 0 0 2-2v-18a2 2 0 0 0-2-2Zm-3.056 11.102H12.056m13.028 5.767H12.056m23.888-11.533H12.056"/>`),
		g.Group(children),
	)
}