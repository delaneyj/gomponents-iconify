package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ngl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.389 18.5v11h5.5m-25.778 0v-11l7.287 11v-11m10.208 3.644a3.644 3.644 0 0 0-3.643-3.644h0a3.644 3.644 0 0 0-3.644 3.644v3.712a3.644 3.644 0 0 0 3.644 3.644h0a3.644 3.644 0 0 0 3.643-3.644h-3.644"/>`),
		g.Group(children),
	)
}