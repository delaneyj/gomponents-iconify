package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PublicEmail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1635 1373l90-90l317 317l-317 317l-90-90l163-163h-668l163 163l-90 90l-317-317l317-317l90 90l-163 163h668l-163-163zM0 384h2048v1024l-128-128V583l-896 449l-896-449v953h640v128H0V384zm271 128l753 376l753-376H271z"/>`),
		g.Group(children),
	)
}