package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterNn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 23h-2v-8h-4v8h-2V13h6a2 2 0 0 1 2 2zm-12-4L9.32 9H7v14h2V13l3.68 10H15V9h-2v10z"/>`),
		g.Group(children),
	)
}