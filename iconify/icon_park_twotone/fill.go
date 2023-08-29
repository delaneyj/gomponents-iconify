package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFill0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path fill="#555" d="M14 14h20v20H14z"/><path d="M34 23L23 34m2-20L14 25m20-11L14 34m0-12v12h12m-4-20h12v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFill0)"/>`),
		g.Group(children),
	)
}