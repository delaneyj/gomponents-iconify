package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2A10 10 0 0 0 2 12a10 10 0 0 0 10 10a10 10 0 0 0 10-10A10 10 0 0 0 12 2m-1 12.41v5.52a7.988 7.988 0 0 1-3.9-1.62l3.9-3.9m2 0l3.9 3.9a7.988 7.988 0 0 1-3.9 1.62v-5.52M4 12c0-4.03 3-7.43 7-7.93v7.52L5.69 16.9A7.913 7.913 0 0 1 4 12m14.31 4.9L13 11.59V4.07c4 .5 7 3.9 7 7.93c0 1.78-.59 3.5-1.69 4.9Z"/>`),
		g.Group(children),
	)
}