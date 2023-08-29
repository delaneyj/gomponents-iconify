package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuUnfoldOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMenuUnfoldOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 10.5h32m-16 9h16m-16 9h16m-32 9h32"/><path fill="#555" d="m16 19l-8 5l8 5V19Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMenuUnfoldOne0)"/>`),
		g.Group(children),
	)
}