package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForChile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM18.385 16.513L20.2 11.35l1.87 5.163h5.996l-4.84 3.352l1.836 5.253l-4.862-3.229l-4.862 3.229l1.836-5.253l-4.84-3.352h6.051zM4 31v-2h28V4c15.104 0 27.471 12.022 28 27H4z"/>`),
		g.Group(children),
	)
}