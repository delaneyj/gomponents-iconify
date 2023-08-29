package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForLatvia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm28 30c0 2.374-.299 4.68-.858 6.883H4.858A27.954 27.954 0 0 1 4 32c0-2.375.299-4.681.858-6.884h54.283c.56 2.203.859 4.509.859 6.884z"/>`),
		g.Group(children),
	)
}