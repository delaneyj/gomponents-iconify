package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Computer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSComputer0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M19 32h10v9H19z"/><rect width="38" height="24" x="5" y="8" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M22 27h4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M14 41h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSComputer0)"/>`),
		g.Group(children),
	)
}