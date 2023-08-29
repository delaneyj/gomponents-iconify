package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M134 120v56a6 6 0 0 1-12 0v-56a6 6 0 0 1 12 0Zm103.88-22.15l-13.88 104A14 14 0 0 1 210.13 214H45.87A14 14 0 0 1 32 201.85l-13.87-104A14 14 0 0 1 32 82h37.28l54.2-61.95a6 6 0 0 1 9 0l54.2 62H224a14 14 0 0 1 13.87 15.85ZM85.22 82h85.56L128 33.11ZM225.5 94.68A2 2 0 0 0 224 94H32a2 2 0 0 0-1.51.68a2 2 0 0 0-.49 1.58l13.86 104a2 2 0 0 0 2 1.73h164.27a2 2 0 0 0 2-1.73l13.87-104a1.93 1.93 0 0 0-.5-1.58ZM181.4 114a6 6 0 0 0-6.57 5.37l-5.6 56a6 6 0 0 0 5.37 6.63h.61a6 6 0 0 0 6-5.4l5.6-56a6 6 0 0 0-5.41-6.6Zm-100.23 5.4a6 6 0 0 0-11.94 1.2l5.6 56a6 6 0 0 0 6 5.4h.61a6 6 0 0 0 5.37-6.57Z"/>`),
		g.Group(children),
	)
}