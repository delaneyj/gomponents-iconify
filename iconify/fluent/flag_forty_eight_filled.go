package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.75 6c-.69 0-1.25.56-1.25 1.25v35.5a1.25 1.25 0 1 0 2.5 0V33h31.25a1.25 1.25 0 0 0 1.007-1.99L33.801 19.5l8.456-11.51A1.25 1.25 0 0 0 41.25 6H8.75Z"/>`),
		g.Group(children),
	)
}