package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.75 0H2.25A2.25 2.25 0 0 0 0 2.25v19.5A2.25 2.25 0 0 0 2.25 24h10.5A2.25 2.25 0 0 0 15 21.75V2.25A2.25 2.25 0 0 0 12.75 0zM7.5 22.5a1.498 1.498 0 1 1 .002-2.996A1.498 1.498 0 0 1 7.5 22.5h-.001z"/>`),
		g.Group(children),
	)
}