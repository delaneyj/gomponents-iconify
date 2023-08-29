package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterRr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 15v-4a2 2 0 0 0-2-2H8v14h2v-6h1.48l2.34 6H16l-2.33-6H14a2 2 0 0 0 2-2zm-6-4h4v4h-4zm14 2h-6v10h2v-8h4v-2z"/>`),
		g.Group(children),
	)
}