package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MacFinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMacFinder0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M44 38V10a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2Z"/><path stroke="#000" d="M25 8s-5 10-4 17h6l1 15"/><path stroke="#fff" d="M34 40H22m8-32H18"/><path stroke="#000" d="M34 16v2m-20-2v2m-1 11s4.19 3 11 3s11-3 11-3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMacFinder0)"/>`),
		g.Group(children),
	)
}