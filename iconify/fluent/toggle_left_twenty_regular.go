package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeftTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm12-2a4 4 0 0 0-4-4H6a4 4 0 1 0 0 8h8a4 4 0 0 0 4-4Zm-4-3a3 3 0 1 1 0 6H6a3 3 0 1 1 0-6h8Z"/>`),
		g.Group(children),
	)
}