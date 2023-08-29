package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendToBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSendToBack0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M14 21H5V5h16v9"/><path stroke-linecap="round" d="M32 27h11v16H27V32"/><path fill="#fff" d="M14 32V14h18v18H14Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSendToBack0)"/>`),
		g.Group(children),
	)
}