package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.74 31.29v3.77a1.64 1.64 0 0 0 1.71 1.57h20.22l6.93 6.24v-6.24h2.19a1.64 1.64 0 0 0 1.71-1.57V16.37a1.64 1.64 0 0 0-1.71-1.58h-3.92"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.18 8.55a1.68 1.68 0 0 0-1.68 1.68h0v19.38a1.68 1.68 0 0 0 1.68 1.68h30a1.69 1.69 0 0 0 1.69-1.67h0V10.24a1.69 1.69 0 0 0-1.68-1.69h-30Zm28.42 3.26l-13.42 9.88l-13.42-9.88"/>`),
		g.Group(children),
	)
}