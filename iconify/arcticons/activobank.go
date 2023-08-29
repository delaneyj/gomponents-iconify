package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Activobank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 32.822V26.92h8.822v5.902m0-5.902l-2.97-11.742h-2.875L13.5 26.92M31.559 24h-4.974m0-8.822h4.457a2.94 2.94 0 0 1 2.94 2.94v2.941A2.94 2.94 0 0 1 31.043 24h.517a2.94 2.94 0 0 1 2.94 2.94v2.941a2.94 2.94 0 0 1-2.94 2.941h-4.974m-.001-17.644v17.644"/>`),
		g.Group(children),
	)
}