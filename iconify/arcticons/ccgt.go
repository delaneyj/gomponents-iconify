package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ccgt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.6 19.3v8.8a2.9 2.9 0 0 1-2.9 2.9h0a2.88 2.88 0 0 1-2.1-.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.6 19.3h0a2.9 2.9 0 0 1 2.9 2.9v1.9a2.9 2.9 0 0 1-2.9 2.9h0a2.9 2.9 0 0 1-2.9-2.9v-1.9a2.9 2.9 0 0 1 2.9-2.9Zm7.2-2.4v8.8a1.54 1.54 0 0 0 1.5 1.5h.4m-3.5-7.9h3.1m-22.8 6.3a3 3 0 0 1-2.6 1.5h0a2.9 2.9 0 0 1-2.9-2.9v-1.9a2.9 2.9 0 0 1 2.9-2.9h0a3.11 3.11 0 0 1 2.6 1.5m8.5 4.7a3 3 0 0 1-2.6 1.5h0a2.9 2.9 0 0 1-2.9-2.9v-1.9a2.9 2.9 0 0 1 2.9-2.9h0a3.11 3.11 0 0 1 2.6 1.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}