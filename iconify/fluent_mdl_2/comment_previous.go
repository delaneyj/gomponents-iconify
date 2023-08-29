package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentPrevious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1024h-128V256H128v1152h256v293l293-293h222l-64 64l64 64H731l-475 475v-475H0V128h2048zm-614 1280h614v128h-614l163 163l-90 90l-318-317l318-317l90 90l-163 163z"/>`),
		g.Group(children),
	)
}