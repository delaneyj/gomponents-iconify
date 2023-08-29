package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersHorizontalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M40 84h36.29a28 28 0 0 0 55.42 0H216a4 4 0 0 0 0-8h-84.29a28 28 0 0 0-55.42 0H40a4 4 0 0 0 0 8Zm64-24a20 20 0 1 1-20 20a20 20 0 0 1 20-20Zm112 112h-20.29a28 28 0 0 0-55.42 0H40a4 4 0 0 0 0 8h100.29a28 28 0 0 0 55.42 0H216a4 4 0 0 0 0-8Zm-48 24a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}