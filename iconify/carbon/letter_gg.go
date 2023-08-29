package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterGg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 13a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h4v2h-5v2h5a2 2 0 0 0 2-2V13zm4 8h-4v-6h4zm-8 2H9a2 2 0 0 1-2-2V11a2 2 0 0 1 2-2h6v2H9v10h4v-4h-2v-2h4z"/>`),
		g.Group(children),
	)
}