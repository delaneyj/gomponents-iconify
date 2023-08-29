package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSendReplyEmailReplyMessageActionsActionArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 12.5C10.58 10 9.86 9 7 9H5.5v3l-5-5.5l5-5v3h1c5 0 6 5 7 8Z"/>`),
		g.Group(children),
	)
}