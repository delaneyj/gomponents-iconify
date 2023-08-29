package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.45 15.6L7.85 5H15q.5 0 .925.213t.725.637l3.525 5q.375.525.375 1.15t-.375 1.15L18.45 15.6ZM5 19q-.825 0-1.413-.588T3 17V7q0-.275.063-.5t.187-.45L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-3.15-3.15q-.225.125-.45.188T15 19H5Z"/>`),
		g.Group(children),
	)
}