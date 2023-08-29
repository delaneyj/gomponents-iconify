package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hammer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#a57939" d="m37.877 21.43l-.029 22.122l1.202 3.623l-.025 15.706l-1.208.002l-3.624.006l-1.209.002l.025-15.706l1.214-3.627l.029-22.122l3.625-.006z"/><path fill="#9b9b9a" d="M29.703 11.6v9.9h12.728v-5.658s8.486 0 11.314 2.829A15.144 15.144 0 0 0 42.431 11.6a60.523 60.523 0 0 0-12.727 0Z"/><path fill="#d0cfce" d="m25.89 12.479l-.032 8.02a3.336 3.336 0 0 1-3.701.014l.032-8.02a2.617 2.617 0 0 1 3.701-.014Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m37.877 21.5l-.029 22.052l1.202 3.623l-.025 15.706l-1.208.002l-3.624.006l-1.209.002l.025-15.706l1.214-3.627l.029-22.058h3.625zm-8.174-9.9v9.9h12.728v-5.658s8.486 0 11.314 2.829A15.144 15.144 0 0 0 42.431 11.6a60.523 60.523 0 0 0-12.727 0Zm-3.813.879l-.032 8.02h0a3.336 3.336 0 0 1-3.701.014h0l.032-8.02h0a2.617 2.617 0 0 1 3.701-.014Z"/>`),
		g.Group(children),
	)
}