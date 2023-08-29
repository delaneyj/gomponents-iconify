package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterJj(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 9h2v2h-2zm0 16h-3v2h3a2 2 0 0 0 2-2V13h-2zm-6-2h-4a2 2 0 0 1-2-2v-2h2v2h4V9h2v12a2 2 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}