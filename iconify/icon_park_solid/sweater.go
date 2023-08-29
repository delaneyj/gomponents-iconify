package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sweater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSweater0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M14 37H6V24l5-15l8-5h10l9 5l4 15v13h-8v7H14v-7Z"/><path d="M34 28v9m-20-9v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSweater0)"/>`),
		g.Group(children),
	)
}