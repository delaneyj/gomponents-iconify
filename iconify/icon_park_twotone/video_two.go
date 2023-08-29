package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVideoTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path fill="#555" d="M20.5 28v-6.062l5.25 3.03L31 28l-5.25 3.031l-5.25 3.031V28Z"/><path d="M6 15h36m-9-9l-6 9m-6-9l-6 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVideoTwo0)"/>`),
		g.Group(children),
	)
}