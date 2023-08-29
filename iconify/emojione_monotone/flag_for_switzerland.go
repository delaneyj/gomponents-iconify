package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSwitzerland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm14.75 34.917h-9.833v9.833h-9.834v-9.833H17.25v-9.834h9.833V17.25h9.834v9.833h9.833v9.834z"/>`),
		g.Group(children),
	)
}