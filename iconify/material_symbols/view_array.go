package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewArray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 19V5h3v14h-3ZM7 19V5h10v14H7Zm-4 0V5h3v14H3Z"/>`),
		g.Group(children),
	)
}