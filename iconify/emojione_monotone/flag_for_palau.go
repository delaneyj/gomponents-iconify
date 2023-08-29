package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForPalau(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-7.867 45.733C15.446 47.733 8.4 40.687 8.4 32s7.046-15.733 15.732-15.733c8.688 0 15.734 7.047 15.734 15.733S32.82 47.733 24.133 47.733z"/>`),
		g.Group(children),
	)
}