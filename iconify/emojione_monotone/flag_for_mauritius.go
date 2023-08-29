package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMauritius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 2c9.61 0 18.104 4.868 23.148 12.267H8.852C13.896 8.868 22.39 4 32 4zm23.148 43.733H8.852A27.834 27.834 0 0 1 4 32h56a27.834 27.834 0 0 1-4.852 15.733z"/>`),
		g.Group(children),
	)
}