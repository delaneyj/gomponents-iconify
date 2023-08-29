package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkKnifeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M76 88V40a4 4 0 0 1 8 0v48a4 4 0 0 1-8 0Zm136-48v184a4 4 0 0 1-8 0v-52h-52a4 4 0 0 1-4-4a264.27 264.27 0 0 1 7.11-55.94c9.47-39.22 27.21-65.41 51.31-75.74A4 4 0 0 1 212 40Zm-8 6.46c-41.75 23.87-47.19 99.29-47.9 117.54H204Zm-88-7.12a4 4 0 0 0-7.9 1.32l8 47.66a36 36 0 0 1-72 0l8-47.66a4 4 0 0 0-7.9-1.32l-8 48a4.89 4.89 0 0 0-.2.66a44.06 44.06 0 0 0 40 43.81V224a4 4 0 0 0 8 0v-92.19A44.06 44.06 0 0 0 124 88a4.89 4.89 0 0 0 0-.66Z"/>`),
		g.Group(children),
	)
}