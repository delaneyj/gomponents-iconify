package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInbox0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4 30L9 6h30l5 24"/><path fill="#555" d="M4 30h10.91l1.817 6h14.546l1.818-6H44v13H4V30Z"/><path stroke-linecap="round" d="M19 14h10m-13 8h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInbox0)"/>`),
		g.Group(children),
	)
}