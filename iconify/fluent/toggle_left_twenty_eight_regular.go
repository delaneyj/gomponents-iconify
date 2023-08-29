package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeftTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 17a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM8 8a6 6 0 0 0 0 12h12a6 6 0 0 0 0-12H8Zm-4.5 6A4.5 4.5 0 0 1 8 9.5h12a4.5 4.5 0 1 1 0 9H8A4.5 4.5 0 0 1 3.5 14Z"/>`),
		g.Group(children),
	)
}