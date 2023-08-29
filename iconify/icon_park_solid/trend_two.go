package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTrendTwo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4 44h40"/><path fill="#fff" d="m4 26l8 2v10H4V26Zm16-2l8-4v18h-8V24Zm16-8l8-4v26h-8V16Z"/><path stroke-linecap="round" d="m4 18l8 2L44 4H34"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTrendTwo0)"/>`),
		g.Group(children),
	)
}