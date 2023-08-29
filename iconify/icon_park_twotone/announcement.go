package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Announcement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAnnouncement0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="40" height="26" x="4" y="15" fill="#555" rx="2"/><path fill="#555" stroke-linecap="round" d="m24 7l-8 8h16l-8-8Z"/><path stroke-linecap="round" d="M12 24h18m-18 8h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAnnouncement0)"/>`),
		g.Group(children),
	)
}