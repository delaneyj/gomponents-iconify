package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnifeFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 4v40M8 5v10c0 5 6 5 6 5s6 0 6-5V5"/><path d="M30 12c0-8 8-8 8-8v17h-8v-9Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 21h-8v-9c0-8 8-8 8-8v17Zm0 0v23"/></g>`),
		g.Group(children),
	)
}