package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSendForwardEmailEmailSendMessageEnvelopeActionsActionForwardArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 12.5C3.42 10 4.14 9 7 9h1.5v3l5-5.5l-5-5v3h-1c-5 0-6 5-7 8Z"/>`),
		g.Group(children),
	)
}