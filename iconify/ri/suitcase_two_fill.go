package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 23h-2v-1H8v1H6v-1H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h3V3a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2h3a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2h-1v1ZM10 9H8v9h2V9Zm6 0h-2v9h2V9Zm-2-5h-4v1h4V4Z"/>`),
		g.Group(children),
	)
}