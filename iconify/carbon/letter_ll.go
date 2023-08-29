package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 21V9H9v14h8v-2h-6zm12 2h-2a2 2 0 0 1-2-2V9h2v12h2z"/>`),
		g.Group(children),
	)
}