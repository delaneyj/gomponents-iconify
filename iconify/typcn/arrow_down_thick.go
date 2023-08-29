package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownThick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.414 10.656a2 2 0 0 0-2.828 0L14 12.242V5a2 2 0 0 0-4 0v7.242l-1.586-1.586a2 2 0 1 0-2.828 2.828L12 19.898l6.414-6.414a2 2 0 0 0 0-2.828z"/>`),
		g.Group(children),
	)
}