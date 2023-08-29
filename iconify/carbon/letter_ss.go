package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterSs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 23h-5v-2h5v-2h-3a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2h5v2h-5v2h3a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2zm-9 0H7v-2h6v-4H9a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h6v2H9v4h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}