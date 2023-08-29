package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxDownloadR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInboxDownloadR0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path stroke-linecap="round" d="M4 31h11l2 4h14l2-4h11m-2 5V26m-24-4l6 6l6-6m-6 6V12M6 36V26"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInboxDownloadR0)"/>`),
		g.Group(children),
	)
}