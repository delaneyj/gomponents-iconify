package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterFf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 11V9H8v14h2v-6h5v-2h-5v-4h6zm8 0V9h-3a2 2 0 0 0-2 2v2h-2v2h2v8h2v-8h3v-2h-3v-2z"/>`),
		g.Group(children),
	)
}