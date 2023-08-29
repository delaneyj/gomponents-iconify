package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm3.198 33.627v10.881h-6.063V35.627L19.096 17.492h7.146l6.023 12.637l5.769-12.637h6.87l-9.706 18.135z"/>`),
		g.Group(children),
	)
}