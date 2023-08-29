package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrequencyGenerator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Zm-29.6 9.6v17.8m11.8-17.8v17.8M10.9 24h11.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.2 21.1h8.9l-8.9 11.8h8.9"/>`),
		g.Group(children),
	)
}