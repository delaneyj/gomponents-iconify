package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19V5h5.325v14H3Zm6.325 0V5h5.325v14H9.325Zm6.325 0V5h5.325v14H15.65Z"/>`),
		g.Group(children),
	)
}