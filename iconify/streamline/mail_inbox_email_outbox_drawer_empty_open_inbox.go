package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailInboxEmailOutboxDrawerEmptyOpenInbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 8H4a1 1 0 0 1 1 1a2 2 0 0 0 4 0a1 1 0 0 1 1-1h3.5"/></g>`),
		g.Group(children),
	)
}