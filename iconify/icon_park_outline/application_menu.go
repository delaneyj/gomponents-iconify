package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationMenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M10 14a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM10 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM10 42a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`),
		g.Group(children),
	)
}