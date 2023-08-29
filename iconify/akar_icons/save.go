package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10a4 4 0 0 0 4-4V7.414a1 1 0 0 0-.293-.707l-3.414-3.414A1 1 0 0 0 16.586 3H7a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4Z"/><path d="M9 3h6v3a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1V3Zm8 18v-7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v7"/><path stroke-linecap="round" d="M11 17h2"/></g>`),
		g.Group(children),
	)
}