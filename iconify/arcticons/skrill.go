package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skrill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.784 29.288v-4.13a2.503 2.503 0 0 1 2.503-2.503h0m-9.907-3.38v10.013m.001-2.127l4.533-4.51m-3.09 3.075l3.563 3.547"/><circle cx="30.556" cy="19.588" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.556 22.654v6.634M9.698 28.191a2.8 2.8 0 0 0 2.455 1.097h1.483a2.5 2.5 0 0 0 2.498-2.503h0a2.5 2.5 0 0 0-2.498-2.503h-1.638A2.5 2.5 0 0 1 9.5 21.778h0a2.5 2.5 0 0 1 2.498-2.503h1.482a2.8 2.8 0 0 1 2.455 1.097m17.218-1.097v8.762a1.252 1.252 0 0 0 1.252 1.251h.375m2.093-10.013v8.762a1.252 1.252 0 0 0 1.252 1.251h.375"/>`),
		g.Group(children),
	)
}