package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusSquareO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 7H9V4H7v3H4v2h3v3h2V9h3z"/><path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/>`),
		g.Group(children),
	)
}