package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenSizeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.975 19.025Q4.75 18.8 4.75 18.5t.225-.525l13-13q.225-.225.525-.225t.525.225q.225.225.225.525t-.225.525l-13 13q-.225.225-.525.225t-.525-.225Z"/>`),
		g.Group(children),
	)
}