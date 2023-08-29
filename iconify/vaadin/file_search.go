package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 13.47V15H2V1h6v4h4v.56c.386.229.716.504.996.825L13 4L9 0H1v16h12v-1.53zM9 1l3 3H9V1z"/><path fill="currentColor" d="m14.78 12.72l-1.92-1.92a.727.727 0 0 0-.325-.179a3.014 3.014 0 0 0 .468-1.618a3 3 0 1 0-1.371 2.52c.02.136.083.248.169.337l1.92 1.92a.75.75 0 0 0 1.059-1.061zM10 11a2 2 0 1 1-.001-3.999A2 2 0 0 1 10 11z"/>`),
		g.Group(children),
	)
}