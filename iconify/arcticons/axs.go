package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM24.19 24l-3.31 4.38m6.62-8.76l-1.76 2.33m3.96 9.35l-8.82-11.68"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.81 27.64a3.72 3.72 0 0 0 2.72.74h.75a2.19 2.19 0 0 0 2.18-2.19h0A2.19 2.19 0 0 0 35.28 24h-1.49a2.18 2.18 0 0 1-2.18-2.19h0a2.18 2.18 0 0 1 2.18-2.19h.74a3.73 3.73 0 0 1 2.73.74m-20.11 4.72a3.31 3.31 0 0 1-3.31 3.3h0a3.3 3.3 0 0 1-3.3-3.3v-2.16a3.3 3.3 0 0 1 3.3-3.3h0a3.31 3.31 0 0 1 3.31 3.3m0 5.46v-8.76"/>`),
		g.Group(children),
	)
}