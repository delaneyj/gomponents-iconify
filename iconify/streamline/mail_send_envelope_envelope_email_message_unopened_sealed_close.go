package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSendEnvelopeEnvelopeEmailMessageUnopenedSealedClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="10.5" x=".5" y="1.75" rx="1"/><path d="m.5 3l5.86 5a1 1 0 0 0 1.28 0l5.86-5"/></g>`),
		g.Group(children),
	)
}