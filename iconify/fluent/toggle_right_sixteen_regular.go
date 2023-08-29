package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleRightSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM1 8a4 4 0 0 1 4-4h6a4 4 0 0 1 0 8H5a4 4 0 0 1-4-4Zm4-3a3 3 0 0 0 0 6h6a3 3 0 1 0 0-6H5Z"/>`),
		g.Group(children),
	)
}