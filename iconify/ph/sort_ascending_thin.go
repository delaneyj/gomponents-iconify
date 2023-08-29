package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAscendingThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M124 128a4 4 0 0 1-4 4H48a4 4 0 0 1 0-8h72a4 4 0 0 1 4 4ZM48 68h136a4 4 0 0 0 0-8H48a4 4 0 0 0 0 8Zm56 120H48a4 4 0 0 0 0 8h56a4 4 0 0 0 0-8Zm122.83-22.83a4 4 0 0 0-5.66 0L188 198.34V112a4 4 0 0 0-8 0v86.34l-33.17-33.17a4 4 0 0 0-5.66 5.66l40 40a4 4 0 0 0 5.66 0l40-40a4 4 0 0 0 0-5.66Z"/>`),
		g.Group(children),
	)
}