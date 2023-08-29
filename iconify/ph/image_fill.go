package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H40a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm-60 48a12 12 0 1 1-12 12a12 12 0 0 1 12-12ZM40 200v-28l52-52l80 80Zm176 0h-21.37l-36-36l20-20L216 181.38V200Z"/>`),
		g.Group(children),
	)
}