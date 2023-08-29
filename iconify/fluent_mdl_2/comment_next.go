package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1027l-128-128V256H128v1152h256v293l293-293h219v128H731l-475 475v-475H0V128h2048zm-355 1027l318 317l-318 317l-90-90l163-163h-614v-128h614l-163-163l90-90z"/>`),
		g.Group(children),
	)
}