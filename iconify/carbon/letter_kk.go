package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterKk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 9h-2.11L9 15.55V9H7v14h2v-4.29l.93-1.49L12.89 23H15l-3.89-7.57L15 9zm7.78 14H25l-3.78-6L25 13h-2.24L19 17.17V9h-2v14h2v-3.75l.96-1.04L22.78 23z"/>`),
		g.Group(children),
	)
}