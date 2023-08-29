package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSInboxR0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" d="M4 31h11l2 4h14l2-4h11"/><path stroke="#fff" stroke-linecap="round" d="M42 36V26M6 36V26"/><path stroke="#000" stroke-linecap="round" d="M17 15h14m-14 8h14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSInboxR0)"/>`),
		g.Group(children),
	)
}