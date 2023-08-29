package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FadersHorizontalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M34 80a6 6 0 0 1 6-6h32a6 6 0 0 1 0 12H40a6 6 0 0 1-6-6Zm182 90h-42v-18a6 6 0 0 0-12 0v48a6 6 0 0 0 12 0v-18h42a6 6 0 0 0 0-12Zm-80 0H40a6 6 0 0 0 0 12h96a6 6 0 0 0 0-12Zm-32-60a6 6 0 0 0 6-6V86h106a6 6 0 0 0 0-12H110V56a6 6 0 0 0-12 0v48a6 6 0 0 0 6 6Z"/>`),
		g.Group(children),
	)
}