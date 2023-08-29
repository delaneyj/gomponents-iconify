package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.51 19a5 5 0 1 0 5 5a5 5 0 0 0-5-5ZM9 28.82v-9.6l7.39 9.6v-9.6m15.92 8.69A2.76 2.76 0 0 0 34.74 29h1.78A2.48 2.48 0 0 0 39 26.51h0A2.48 2.48 0 0 0 36.52 24h-1.93a2.48 2.48 0 0 1-2.48-2.48h0a2.48 2.48 0 0 1 2.48-2.48h1.78a2.76 2.76 0 0 1 2.43 1.09"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.51 21.51 0 0 0 2.5 24h0A21.51 21.51 0 0 0 24 45.5h0a21.5 21.5 0 0 0 0-43Z"/>`),
		g.Group(children),
	)
}