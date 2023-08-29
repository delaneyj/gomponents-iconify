package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 104v48a16 16 0 0 1-16 16h-48v48a16 16 0 0 1-16 16h-48a16 16 0 0 1-16-16v-48H40a16 16 0 0 1-16-16v-48a16 16 0 0 1 16-16h48V40a16 16 0 0 1 16-16h48a16 16 0 0 1 16 16v48h48a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}