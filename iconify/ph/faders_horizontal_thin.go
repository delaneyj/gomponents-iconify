package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FadersHorizontalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M36 80a4 4 0 0 1 4-4h32a4 4 0 0 1 0 8H40a4 4 0 0 1-4-4Zm180 92h-44v-20a4 4 0 0 0-8 0v48a4 4 0 0 0 8 0v-20h44a4 4 0 0 0 0-8Zm-80 0H40a4 4 0 0 0 0 8h96a4 4 0 0 0 0-8Zm-32-64a4 4 0 0 0 4-4V84h108a4 4 0 0 0 0-8H108V56a4 4 0 0 0-8 0v48a4 4 0 0 0 4 4Z"/>`),
		g.Group(children),
	)
}