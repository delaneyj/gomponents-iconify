package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterTt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 11h3v12h2V11h3V9H8v2zm15 4v-2h-3v-2h-2v2h-2v2h2v6a2 2 0 0 0 2 2h3v-2h-3v-6z"/>`),
		g.Group(children),
	)
}