package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 52H32a4 4 0 0 0-4 4v136a12 12 0 0 0 12 12h176a12 12 0 0 0 12-12V56a4 4 0 0 0-4-4ZM36 108h48v40H36Zm56 0h128v40H92Zm128-48v40H36V60ZM36 192v-36h48v40H40a4 4 0 0 1-4-4Zm180 4H92v-40h128v36a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}