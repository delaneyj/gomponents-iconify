package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineWbShade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 12v2.5l5.5 5.5H22zm0 8h3l-3-3zM8 4l-6 6h2v10h8V10h2L8 4zm1 10H7v-4h2v4z"/>`),
		g.Group(children),
	)
}