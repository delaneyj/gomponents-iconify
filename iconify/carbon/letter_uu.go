package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterUu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 23h-4a2 2 0 0 1-2-2v-8h2v8h4v-8h2v8a2 2 0 0 1-2 2zm-10 0H9a2 2 0 0 1-2-2V9h2v12h4V9h2v12a2 2 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}