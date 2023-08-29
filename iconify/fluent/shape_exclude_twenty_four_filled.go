package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeExcludeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 2A3.25 3.25 0 0 0 2 5.25v8a3.25 3.25 0 0 0 3.25 3.25H7.5v2.25A3.25 3.25 0 0 0 10.75 22h8A3.25 3.25 0 0 0 22 18.75v-8a3.25 3.25 0 0 0-3.25-3.25H16.5V5.25A3.25 3.25 0 0 0 13.25 2h-8ZM16.5 7.5v5.75a3.25 3.25 0 0 1-3.25 3.25H7.5v-5.75a3.25 3.25 0 0 1 3.25-3.25h5.75Z"/>`),
		g.Group(children),
	)
}