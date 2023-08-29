package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PriceTagThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.904 2.1l9.9 1.414l1.414 9.9l-9.193 9.192a1 1 0 0 1-1.414 0l-9.9-9.9a1 1 0 0 1 0-1.413L10.905 2.1Zm2.829 8.485a2 2 0 1 0 2.828-2.828a2 2 0 0 0-2.828 2.828Z"/>`),
		g.Group(children),
	)
}