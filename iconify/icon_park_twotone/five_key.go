package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFiveKey0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path d="M29 14.01h-9v7.024C20 21 22 20 25 20s4 2.034 4 6s-1 7-5 7c-3 0-4-2-4-3.992"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFiveKey0)"/>`),
		g.Group(children),
	)
}