package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clima(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36 22.766a7.456 7.456 0 0 0-3.754 1.014a10.478 10.478 0 0 0-16.492 0a7.5 7.5 0 1 0 0 12.973a10.478 10.478 0 0 0 16.492 0A7.497 7.497 0 1 0 36 22.766Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.851 20.96q.005-.13.005-.262a7 7 0 1 0-13.352 2.944m6.352-13.408v-3m-7.399 6.064l-2.121-2.121m-.943 9.52h-3m20.862-7.399l2.121-2.121m.943 9.52h3"/>`),
		g.Group(children),
	)
}