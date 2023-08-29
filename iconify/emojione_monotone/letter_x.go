package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m5.513 44.508l-5.514-9.894l-5.825 9.894h-7.048l9.331-14.783l-8.878-14.232h7.244l5.175 9.449l5.317-9.449h7.008l-8.878 13.996l9.429 15.02h-7.361z"/>`),
		g.Group(children),
	)
}