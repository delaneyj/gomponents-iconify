package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sverigesradioplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.11 27.3a3.34 3.34 0 0 0 2.46.67h.67a2 2 0 0 0 2-2h0a2 2 0 0 0-2-2H19.9a2 2 0 0 1-2-2h0a2 2 0 0 1 2-2h.67a3.38 3.38 0 0 1 2.43.73m-7.9 5.76A3 3 0 0 1 12.5 28h0a3 3 0 0 1-3-3v-2a3 3 0 0 1 3-3h0a3 3 0 0 1 3 3v1h-6m29 1a3 3 0 0 1-3 3h0a3 3 0 0 1-3-3v-2a3 3 0 0 1 3-3h0a3 3 0 0 1 3 3m0 4.97v-7.94M25.67 20h2.39a2 2 0 0 1 2 2h0a2 2 0 0 1-2 2h-2.39m0 0l4.38 3.97"/>`),
		g.Group(children),
	)
}