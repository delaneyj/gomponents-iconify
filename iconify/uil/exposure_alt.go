package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposureAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 6H7a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Zm8-4H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3ZM4 18.59V5a1 1 0 0 1 1-1h13.59ZM20 19a1 1 0 0 1-1 1H5.41L20 5.41Zm-7-2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}