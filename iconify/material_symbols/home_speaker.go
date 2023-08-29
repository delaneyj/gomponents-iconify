package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSpeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.2 12l.775-5.825q.05-.275.2-.488t.425-.312l7.9-3.15q.45-.2.875.063t.475.762L17.925 12H6.2Zm3.375 9Q7.75 21 6.55 19.637t-.95-3.162L5.925 14H18.15l.3 2.525q.225 1.8-.962 3.138T14.5 21H9.575Z"/>`),
		g.Group(children),
	)
}