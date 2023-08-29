package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleRightTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 17a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm1-9a6 6 0 0 1 0 12H8A6 6 0 0 1 8 8h12Zm4.5 6A4.5 4.5 0 0 0 20 9.5H8a4.5 4.5 0 1 0 0 9h12a4.5 4.5 0 0 0 4.5-4.5Z"/>`),
		g.Group(children),
	)
}