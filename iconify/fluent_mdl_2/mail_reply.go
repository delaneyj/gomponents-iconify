package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v1082q-29-23-61-39t-67-29V583l-896 449l-896-449v953h1024v128H0V384h2048zM1024 888l753-376H271l753 376zm475 648q50 0 110-2t122 0t118 14t101 40t71 78t27 126q0 53-20 99t-55 81t-82 55t-99 21v-128q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10h-293l162 163l-90 90l-317-317l317-317l90 90l-162 163z"/>`),
		g.Group(children),
	)
}