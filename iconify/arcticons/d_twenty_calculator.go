package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DTwentyCalculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5h33c1.1 0 2 .9 2 2v33c0 1.1-.9 2-2 2h-33c-1.1 0-2-.9-2-2v-33c0-1.1.9-2 2-2Zm16.5 0v37M5.5 24h37m-13.8-8.6h9M11.3 29.3l8 7.9m-.1-7.9l-8 8m4.1-26.4v9m-4.5-4.5h9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.5 28l-2.2 6.8l4.8 5.2l6.9-1.5l2.2-6.8l-4.8-5.2l-6.9 1.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.5 28l1.5 6.2l1.1 5.8l4.6-4.3l4.5-4l-5.9-1.8l-5.8-1.9z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.7 35.7L30 34.2l4.3-4.3l1.4 5.8zm-9.4-.9l3.7-.6m4.3-4.3l1.1-3.4m2.6 12l-2.3-2.8"/>`),
		g.Group(children),
	)
}