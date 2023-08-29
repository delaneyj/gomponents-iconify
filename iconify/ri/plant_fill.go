package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlantFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.998 3v2a7 7 0 0 1-7 7h-1v1h5v7a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-7h5v-3a7 7 0 0 1 7-7h3Zm-15.5-1a7.49 7.49 0 0 1 6.124 3.169A7.962 7.962 0 0 0 9.998 10v1h-.5a7.5 7.5 0 0 1-7.5-7.5V2h3.5Z"/>`),
		g.Group(children),
	)
}