package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutDistributeHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4H4v16h2V4zm14 0h-2v16h2V4zM10 7h6v10H8V7h2zm4 8V9h-4v6h4z"/>`),
		g.Group(children),
	)
}