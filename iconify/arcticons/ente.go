package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ente(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.756 25.867a2.226 2.226 0 0 1-1.956 1.15h0a2.307 2.307 0 0 1-2.3-2.3V23.22a2.307 2.307 0 0 1 2.3-2.3h0a2.307 2.307 0 0 1 2.3 2.3v.805h-4.6m22.655 1.85a2.226 2.226 0 0 1-1.955 1.15h0a2.307 2.307 0 0 1-2.3-2.3V23.23a2.307 2.307 0 0 1 2.3-2.3h0a2.307 2.307 0 0 1 2.3 2.3v.805h-4.6M24 27.021v-3.795a2.307 2.307 0 0 0-2.3-2.3h0a2.307 2.307 0 0 0-2.3 2.3v3.795m-.001-3.795v-2.301m8.165-1.951v8.052m-1.265-6.096h2.415"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}