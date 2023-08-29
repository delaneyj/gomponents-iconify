package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kuda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM13 13.118v21.764M19.002 24H13"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.963 34.882L19.002 24l8.961-10.809"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.975 34.882L17.001 24l5.974-10.809"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.988 34.882L15.001 24l2.987-10.809M32.824 30.53A2.176 2.176 0 1 0 35 32.705h0a2.176 2.176 0 0 0-2.176-2.177Z"/>`),
		g.Group(children),
	)
}