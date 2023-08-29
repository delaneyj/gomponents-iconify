package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTrinidadAndTobago(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2M5.699 22.392a27.72 27.72 0 0 1 1.74-3.823L45.433 56.56a27.738 27.738 0 0 1-3.824 1.741L5.699 22.392M18.568 7.439a27.775 27.775 0 0 1 3.823-1.74L58.3 41.608a27.72 27.72 0 0 1-1.74 3.823L18.568 7.439"/>`),
		g.Group(children),
	)
}