package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BigClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBigClock0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 4C12.954 4 4 12.865 4 23.8V44h40V23.8C44 12.865 35.046 4 24 4Z"/><circle cx="24" cy="24" r="12" fill="#000" stroke="#000"/><path stroke="#000" d="M24 18v6l4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBigClock0)"/>`),
		g.Group(children),
	)
}