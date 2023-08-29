package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCongoBrazzaville(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM8.447 47.113L47.113 8.447a28.225 28.225 0 0 1 8.402 8.382L16.83 55.516a28.209 28.209 0 0 1-8.383-8.403z"/>`),
		g.Group(children),
	)
}