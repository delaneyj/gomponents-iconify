package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailForwardMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384H0v1024l128-128V583l896 449l896-449v953H896v128h1152V384zM1024 888L271 512h1506l-753 376zm-611 485l-90-90L6 1600l317 317l90-90l-163-163h518v-128H250l163-163z"/>`),
		g.Group(children),
	)
}