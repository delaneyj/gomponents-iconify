package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.847 15.754a.5.5 0 0 0 .707 0l6.8-6.8a.5.5 0 0 0-.708-.708L10.2 14.693l-2.847-2.846a.5.5 0 0 0-.707.707l3.2 3.2zM21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5zM21 21H3V3h18v18z"/>`),
		g.Group(children),
	)
}