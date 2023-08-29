package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FTwoKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFTwoKey0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path d="M26 20c0-2.533 2.133-4 4-4s4 1.467 4 4c0 4.56-8 7.947-8 12h8M21 16h-7v16m0-8h7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFTwoKey0)"/>`),
		g.Group(children),
	)
}