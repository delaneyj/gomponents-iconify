package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxUploadR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSInboxUploadR0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" d="M4 31h11l2 4h14l2-4h11"/><path stroke="#fff" stroke-linecap="round" d="M42 36V26"/><path stroke="#000" stroke-linecap="round" d="m18 18l6-6l6 6m-6-6v16"/><path stroke="#fff" stroke-linecap="round" d="M6 36V26"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSInboxUploadR0)"/>`),
		g.Group(children),
	)
}