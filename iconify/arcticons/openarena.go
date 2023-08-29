package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openarena(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.275 41.255L24.284 7.436L12.87 41.255m3.805-11.414H31.47"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.825 28.891c-8.09-2.198-12.674-7.625-10.972-12.99s9.22-9.19 17.998-9.156s16.396 3.917 18.236 9.295s-2.604 10.77-10.64 12.907"/>`),
		g.Group(children),
	)
}