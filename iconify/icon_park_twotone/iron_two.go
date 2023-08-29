package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IronTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTIronTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 40h40l-2-16H20c-8.837 0-16 7.163-16 16Z"/><path d="M16 8h24l2 16m-25 8h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTIronTwo0)"/>`),
		g.Group(children),
	)
}