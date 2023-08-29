package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2h2Zm17 2.086V18.21H6v-6.152l6.59-5.99l2.966 3.461L21 4.086Zm-2 4.828l-3.556 3.557l-3.033-3.538L8 12.943v3.267h11V8.914Z"/>`),
		g.Group(children),
	)
}