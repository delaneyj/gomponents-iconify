package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19v-8.93a2 2 0 0 1 .89-1.664l7-4.666a2 2 0 0 1 2.22 0l7 4.666A2 2 0 0 1 21 10.07V19M3 19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2M3 19l6.75-4.5M21 19l-6.75-4.5M3 10l6.75 4.5M21 10l-6.75 4.5m0 0l-1.14.76a2 2 0 0 1-2.22 0l-1.14-.76"/>`),
		g.Group(children),
	)
}