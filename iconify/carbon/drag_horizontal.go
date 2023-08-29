package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 4v11H5.83l2.58-2.59L7 11l-5 5l5 5l1.41-1.41L5.83 17H12v11h2V4h-2zm13 7l-1.41 1.41L26.17 15H20V4h-2v24h2V17h6.17l-2.58 2.59L25 21l5-5l-5-5z"/>`),
		g.Group(children),
	)
}