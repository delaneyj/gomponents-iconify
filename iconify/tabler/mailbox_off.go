package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailboxOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 21v-6.5a3.5 3.5 0 0 0-7 0V21h18m0-4v-2a4 4 0 0 0-4-4h-2m-4 0H6.5M12 8V3h4l2 2l-2 2h-4m-6 8h1M3 3l18 18"/>`),
		g.Group(children),
	)
}