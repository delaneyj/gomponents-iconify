package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterIi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 11h3v10h-3v2h8v-2h-3V11h3V9h-8v2zm10 2h2v10h-2zm0-4h2v2h-2z"/>`),
		g.Group(children),
	)
}