package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderTransgender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 32h-48a8 8 0 0 0 0 16h28.69L168 76.69l-18.34-18.34a8 8 0 1 0-11.32 11.31L156.69 88l-15.76 15.76a71.94 71.94 0 1 0 11.32 11.31L168 99.33l18.34 18.34a8 8 0 0 0 11.32-11.31L179.31 88L208 59.32V88a8 8 0 0 0 16 0V40a8 8 0 0 0-8-8Zm-80.4 167.63A56 56 0 1 1 152 160a56.08 56.08 0 0 1-16.4 39.63Z"/>`),
		g.Group(children),
	)
}