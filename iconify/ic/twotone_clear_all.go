package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneClearAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 11h14v2H5zm-2 4h14v2H3zm4-8h14v2H7z"/>`),
		g.Group(children),
	)
}