package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSendEmailSendEmailPaperAirplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.82 11L8 13.17a1.1 1.1 0 0 0 1.05.3a1.12 1.12 0 0 0 .81-.74L13.44 2A1.12 1.12 0 0 0 12 .56L1.27 4.14A1.12 1.12 0 0 0 .53 5a1.1 1.1 0 0 0 .3 1l2.74 2.74l-.09 3.47ZM13.12.78L3.57 8.74"/>`),
		g.Group(children),
	)
}