package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Osesfrandroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM26.27 28.55h4.55m-4.55-9.1h4.55M26.27 24h2.96m-2.96-4.55v9.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.58 22.5a3 3 0 0 0-3.2-3a3.15 3.15 0 0 0-2.88 3.17v2.83a3 3 0 0 0 3 3.05h0a3 3 0 0 0 3-3.05v-3m4.5-3.05a2.27 2.27 0 0 0-2.26 2.28h0A2.27 2.27 0 0 0 20 24h.77m.03 0h.77a2.27 2.27 0 0 1 2.27 2.27h0a2.27 2.27 0 0 1-2.27 2.28m2.06-8.33a3.87 3.87 0 0 0-2.83-.77H20"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18 27.78a3.85 3.85 0 0 0 2.82.77h.77m13.1-9.1a2.28 2.28 0 0 0-2.27 2.28h0A2.27 2.27 0 0 0 34.69 24h.77m0 0h.77a2.27 2.27 0 0 1 2.27 2.27h0a2.28 2.28 0 0 1-2.27 2.28m2.05-8.33a3.82 3.82 0 0 0-2.82-.77h-.77m-2.05 8.33a3.84 3.84 0 0 0 2.82.77h.77"/>`),
		g.Group(children),
	)
}