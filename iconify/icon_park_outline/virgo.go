package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Virgo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 17c4 0 11.088-.112 11.959 6.64C42.417 27.192 39.238 32.674 24 43"/><path fill="currentColor" d="M16 11a2 2 0 1 0 4 0h-4ZM4 11a2 2 0 1 0 4 0H4Zm4 0a4 4 0 0 1 4-4V3a8 8 0 0 0-8 8h4Zm4-4a4 4 0 0 1 4 4h4a8 8 0 0 0-8-8v4Z"/><path fill="currentColor" d="M28 11a2 2 0 1 0 4 0h-4Zm-12 0a2 2 0 1 0 4 0h-4Zm4 0a4 4 0 0 1 4-4V3a8 8 0 0 0-8 8h4Zm4-4a4 4 0 0 1 4 4h4a8 8 0 0 0-8-8v4Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 11v18m12-18v18"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M30 11v18c0 5 2.5 10 12 10"/></g>`),
		g.Group(children),
	)
}