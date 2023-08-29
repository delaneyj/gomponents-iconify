package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Behance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.005 29.651a3.778 3.778 0 0 1-3.284 1.907h0a3.78 3.78 0 0 1-3.78-3.779v-2.456a3.78 3.78 0 0 1 3.78-3.78h0a3.78 3.78 0 0 1 3.779 3.78v1.228h-7.558"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M19.736 24a3.78 3.78 0 0 1 0 7.559H13.5V16.442h6.236a3.78 3.78 0 0 1 0 7.558Zm0 0H13.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.942 16.442H34.5"/>`),
		g.Group(children),
	)
}