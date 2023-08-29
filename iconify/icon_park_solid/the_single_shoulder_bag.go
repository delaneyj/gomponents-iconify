package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TheSingleShoulderBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTheSingleShoulderBag0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M28 27c0-8.819-1.22-23-4-23c-3.429 0-4 14.181-4 23m-5 0h18v17H15z"/><path fill="#fff" d="M15 27h18l-5.294 9h-7.941L15 27Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTheSingleShoulderBag0)"/>`),
		g.Group(children),
	)
}