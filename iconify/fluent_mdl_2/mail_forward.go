package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384h2048v1024l-128-128V583l-896 449l-896-449v953h1024v128H0V384zm1024 504l753-376H271l753 376zm611 485l90-90l317 317l-317 317l-90-90l163-163h-518v-128h518l-163-163z"/>`),
		g.Group(children),
	)
}