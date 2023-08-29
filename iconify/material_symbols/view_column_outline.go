package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumnOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19V5h17.975v14H3Zm2-2h3.325V7H5v10Zm5.325 0h3.325V7h-3.325v10Zm5.325 0h3.325V7H15.65v10Z"/>`),
		g.Group(children),
	)
}