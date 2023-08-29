package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HlSevenAttributes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M29 9h-8v4h2v-2h3.847L22 23h2.157L29 11V9z" fill="currentColor"/><path d="M14 21V9h-2v14h8v-2h-6z" fill="currentColor"/><path d="M8 9v6H4V9H2v14h2v-6h4v6h2V9H8z" fill="currentColor"/>`),
		g.Group(children),
	)
}