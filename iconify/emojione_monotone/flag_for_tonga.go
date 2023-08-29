package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTonga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M17.25 13.316h-3.934v4.917H8.4v3.934h4.916v4.916h3.934v-4.916h4.917v-3.934H17.25z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-5.9 2.631V32H4C4 18.585 13.484 7.349 26.1 4.631z"/>`),
		g.Group(children),
	)
}