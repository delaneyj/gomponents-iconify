package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Split(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSplit0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 10V7a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v22a2 2 0 0 1-2 2h-4"/><rect width="24" height="24" x="5" y="18" fill="#fff" rx="2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSplit0)"/>`),
		g.Group(children),
	)
}