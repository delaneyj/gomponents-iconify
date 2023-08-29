package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareDownRightF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.828 13.828a.997.997 0 0 0 1-1v-6a1 1 0 1 0-2 0v3.586l-3.95-3.95A1 1 0 0 0 6.465 7.88l3.95 3.95H6.828a1 1 0 0 0 0 2h6zM4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4z"/>`),
		g.Group(children),
	)
}