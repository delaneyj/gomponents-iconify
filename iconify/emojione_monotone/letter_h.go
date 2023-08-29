package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm11.664 44.508h-6.023V33.555h-11.28v12.953h-6.025V17.492h6.025v11.063H37.64V17.492h6.023v29.016z"/>`),
		g.Group(children),
	)
}