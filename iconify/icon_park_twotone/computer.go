package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Computer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTComputer0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M19 32h10v9H19z"/><rect width="38" height="24" x="5" y="8" fill="#555" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M22 27h4M14 41h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTComputer0)"/>`),
		g.Group(children),
	)
}