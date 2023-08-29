package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterEe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 19v-4a2.002 2.002 0 0 0-2-2h-4a2.002 2.002 0 0 0-2 2v6a2.008 2.008 0 0 0 2 2h5v-2h-5v-2zm-6-4h4v2h-4zm-4-4V9H7v14h8v-2H9v-4h5v-2H9v-4h6z"/>`),
		g.Group(children),
	)
}