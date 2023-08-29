package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM14.145 26.7L21.3 37.5m0-10.8l-7.155 10.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.7 17.723a3.578 3.578 0 0 0 7.155 0v-3.645a3.578 3.578 0 1 0-7.155 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M18.6 15.9a2.7 2.7 0 0 1 0 5.4h-4.455V10.5H18.6a2.7 2.7 0 0 1 0 5.4Zm0 0h-4.455"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.7 32.1h7.155"/>`),
		g.Group(children),
	)
}