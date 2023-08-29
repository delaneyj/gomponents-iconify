package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandleA0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#555"/><path d="M24 11L14 33m4-7h12m-6-15l10 22"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandleA0)"/>`),
		g.Group(children),
	)
}