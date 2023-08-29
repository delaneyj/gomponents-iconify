package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 2048v-128h1152v128H384zm1197-979l-621 626l-621-626l90-90l467 470V0h128v1449l467-470l90 90z"/>`),
		g.Group(children),
	)
}