package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1v14h16V1H0zm1 3h6.5v10H1V4zm14 10H8.5V4H15v10zm0-11h-1V2h1v1z"/>`),
		g.Group(children),
	)
}