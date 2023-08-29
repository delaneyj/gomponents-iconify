package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TailoringTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M42 5h-4.586c-.89 0-1.337 1.077-.707 1.707l4.586 4.586c.63.63 1.707.184 1.707-.707V6a1 1 0 0 0-1-1ZM8 43h4.586c.89 0 1.337-1.077.707-1.707l-4.586-4.586c-.63-.63-1.707-.184-1.707.707V42a1 1 0 0 0 1 1Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 4v28a2 2 0 0 0 2 2h28"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 13h14v14"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M14 13H6m30 29v-8"/></g>`),
		g.Group(children),
	)
}