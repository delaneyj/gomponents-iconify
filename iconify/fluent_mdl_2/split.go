package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Split(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1683 595l365 365l-365 365l-91-90l211-211h-651V896h651l-211-211l91-90zM456 685L245 896h651v128H245l211 211l-91 90L0 960l365-365l91 90z"/>`),
		g.Group(children),
	)
}