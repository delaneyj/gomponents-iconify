package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeftSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm10-2a4 4 0 0 0-4-4H5a4 4 0 1 0 0 8h6a4 4 0 0 0 4-4Zm-4-3a3 3 0 1 1 0 6H5a3 3 0 0 1 0-6h6Z"/>`),
		g.Group(children),
	)
}