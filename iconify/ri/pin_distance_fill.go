package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinDistanceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.39 10.39L7.5 14.277L3.61 10.39a5.5 5.5 0 1 1 7.78 0ZM7.5 8.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm12.89 10.89l-3.89 3.888l-3.89-3.889a5.5 5.5 0 1 1 7.78 0ZM16.5 17.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}