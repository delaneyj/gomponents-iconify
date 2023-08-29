package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attapoll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.04 5.5h17.952l9.485 9.485v27.477h-8.462v-8.736H14.18V42.5H5.523V15.002Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.411 14.506V25.23H33.59V14.579Z"/>`),
		g.Group(children),
	)
}