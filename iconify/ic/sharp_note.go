package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 10l-6-6H2v16h20V10zm-7-4.5l5.5 5.5H15V5.5z"/>`),
		g.Group(children),
	)
}