package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M2 6a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v5H2V6Z"/><path d="M18.5 16.5h-3"/><path stroke-linejoin="round" d="M12 11h10v7a4 4 0 0 1-4 4h-6V11Zm0 0H2v7a4 4 0 0 0 4 4h6V11Z"/><path d="M5.5 18L7 16.5m0 0L8.5 15M7 16.5L8.5 18M7 16.5L5.5 15"/></g>`),
		g.Group(children),
	)
}