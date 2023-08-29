package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReplyAllMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384v1082q29-23 61-39t67-29V583l896 449l896-449v953h-704v128h832V384H0zm1024 504L271 512h1506l-753 376zm-91 712l-226 227l90 90l318-317l-318-317l-90 90l226 227zm-384-64q-50 0-110-2t-122 0t-118 14t-101 40t-71 78t-27 126q0 53 20 99t55 81t82 55t99 21v-128q-27 0-50-10t-40-27t-28-41t-10-50q0-27 10-50t27-40t41-28t50-10h293l-162 163l90 90l317-317l-317-317l-90 90l162 163z"/>`),
		g.Group(children),
	)
}