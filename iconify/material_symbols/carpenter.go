package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carpenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.475 21.3q-.275.275-.637.425t-.763.15q-.4 0-.775-.15t-.65-.425l-1.4-1.4q-.275-.275-.412-.612t-.163-.688q-.025-.35.088-.7t.337-.65l.15-.2L3.1 5.4L7 1.5l12.725 12.725q.275.275.425.638t.15.762q0 .4-.15.775t-.425.65l-4.25 4.25Zm-1.4-1.425L18.3 15.65l-1.4-1.425l-4.25 4.25l1.425 1.4Z"/>`),
		g.Group(children),
	)
}