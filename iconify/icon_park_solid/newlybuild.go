package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newlybuild(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNewlybuild0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M19 10V7a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v22a2 2 0 0 1-2 2h-4"/><rect width="24" height="24" x="5" y="18" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" d="M17 25v10m-5-5h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNewlybuild0)"/>`),
		g.Group(children),
	)
}