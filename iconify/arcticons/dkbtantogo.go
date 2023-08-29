package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dkbtantogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.7 38.4V27.2m0 0h4.7a2.8 2.8 0 0 1 2.8 2.8h0a2.8 2.8 0 0 1-2.8 2.8h-4.7m0 0h4.7a2.8 2.8 0 0 1 2.8 2.8h0a2.8 2.8 0 0 1-2.8 2.8h-4.7m-18.9 0V27.2h1.9a5.59 5.59 0 0 1 5.6 5.6h0a5.59 5.59 0 0 1-5.6 5.6h-1.9Zm10.4-11.2v11.2m1.5-5.6l4.1-5.5m-4.1 5.5l4.1 5.6m-4.1-5.6h-1.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.9 14.3h-.6v-2.1a3.8 3.8 0 0 0-7.6 0h0v2.1h-.6a1.37 1.37 0 0 0-1.4 1.4V23a1.37 1.37 0 0 0 1.4 1.4h8.8a1.37 1.37 0 0 0 1.4-1.4h0v-7.3a1.5 1.5 0 0 0-1.4-1.4Zm-8.2 0h7.5"/>`),
		g.Group(children),
	)
}