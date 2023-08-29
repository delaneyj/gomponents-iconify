package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 16h-1V7a1 1 0 0 0-1-1H8V5a1 1 0 0 0-2 0v1H5a1 1 0 0 0 0 2h1v9a1 1 0 0 0 1 1h9v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2Zm-3 0H8V8h8Z"/>`),
		g.Group(children),
	)
}