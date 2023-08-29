package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1v14h16V1H0zm14 1h1v1h-1V2zm1 2v4.5H1V4h14zM1 14V9.5h14V14H1z"/>`),
		g.Group(children),
	)
}