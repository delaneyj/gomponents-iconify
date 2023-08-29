package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorBasicScientificCalculations(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm1.08.945L6.032 41.858m.472-35.413l35.354 35.413"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.253 27a2.658 2.658 0 0 1-2.1 1a2.689 2.689 0 0 1-2.7-2.7v-2.6a2.689 2.689 0 0 1 2.7-2.7a2.814 2.814 0 0 1 2.1 1m-6 1.9h3.4m-3.4 2h3.4" class="b"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.181 33.395h4M22 12.402h4m-2-2v4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.67 27.252a2.384 2.384 0 0 0 2 .9h1.2a2 2 0 0 0 0-4h-1.3a2 2 0 0 1 0-4h1.2a2.147 2.147 0 0 1 2 .9m-2.6-.9v-1m0 10v-1" class="b"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.181 35.757h4"/>`),
		g.Group(children),
	)
}