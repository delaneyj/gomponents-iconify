package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleMultipleTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 8.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-6.5h11a4.5 4.5 0 1 1 0 9h-11a4.5 4.5 0 0 1 0-9Zm11 1.5h-11a3 3 0 0 0 0 6h11a3 3 0 1 0 0-6Zm0 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM2 17.5A4.5 4.5 0 0 1 6.5 13h11a4.5 4.5 0 1 1 0 9h-11A4.5 4.5 0 0 1 2 17.5Zm4.5-3a3 3 0 1 0 0 6h11a3 3 0 1 0 0-6h-11Z"/>`),
		g.Group(children),
	)
}