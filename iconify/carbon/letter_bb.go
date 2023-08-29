package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterBb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 13h-4V9h-2v14h6a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2zm-4 8v-6h4v6zm-4-9a3 3 0 0 0-3-3H7v14h5a3 3 0 0 0 3-3v-2a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2zm-6-1h3a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H9zm4 9a1 1 0 0 1-1 1H9v-4h3a1 1 0 0 1 1 1z"/>`),
		g.Group(children),
	)
}