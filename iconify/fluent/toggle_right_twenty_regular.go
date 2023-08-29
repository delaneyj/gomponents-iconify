package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM2 10a4 4 0 0 1 4-4h8a4 4 0 0 1 0 8H6a4 4 0 0 1-4-4Zm4-3a3 3 0 0 0 0 6h8a3 3 0 1 0 0-6H6Z"/>`),
		g.Group(children),
	)
}