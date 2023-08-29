package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCircleRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-3 9a8.959 8.959 0 0 1-4.49-1.198l-.004-.003c-.128-.073-.192-.11-.253-.127a.475.475 0 0 0-.167-.017c-.064.004-.13.026-.262.07l-2.306.769l-.006.002c-.485.162-.727.242-.889.185a.502.502 0 0 1-.304-.304c-.057-.162.024-.405.186-.892v-.003l.77-2.306c.044-.132.065-.198.07-.261a.498.498 0 0 0-.017-.168a1.228 1.228 0 0 0-.127-.252l-.003-.005A9 9 0 1 1 12 21Z"/>`),
		g.Group(children),
	)
}