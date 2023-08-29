package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M36 16V8H4v24h8"/><path d="M12 40h32V16H12v24Z"/><path stroke-linecap="round" d="m12 16l16 12l16-12"/><path stroke-linecap="round" d="M32 16H12v15"/><path stroke-linecap="round" d="M44 31V16H24"/></g>`),
		g.Group(children),
	)
}