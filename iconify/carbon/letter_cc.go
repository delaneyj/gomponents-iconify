package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterCc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 23h-5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h5v2h-5v6h5zm-9 0H9a2 2 0 0 1-2-2V11a2 2 0 0 1 2-2h6v2H9v10h6z"/>`),
		g.Group(children),
	)
}