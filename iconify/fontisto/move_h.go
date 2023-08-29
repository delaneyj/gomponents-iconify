package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.781 14.819V9.6H24v5.219zm-9.391 0V9.6h5.219v5.219zm-9.39 0V9.6h5.219v5.219z"/>`),
		g.Group(children),
	)
}