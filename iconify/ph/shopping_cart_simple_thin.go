package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 216a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm92-12a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm43.82-130.82l-28.52 92.7A19.9 19.9 0 0 1 180.18 180H84.07a20.08 20.08 0 0 1-19.23-14.51L28.67 38.9a4 4 0 0 0-3.85-2.9H8a4 4 0 0 1 0-8h16.82a12.05 12.05 0 0 1 11.54 8.7L45.3 68H224a4 4 0 0 1 3.82 5.18ZM218.58 76h-171l24.94 87.3a12.05 12.05 0 0 0 11.55 8.7h96.11a11.94 11.94 0 0 0 11.47-8.47Z"/>`),
		g.Group(children),
	)
}