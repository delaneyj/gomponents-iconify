package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatMinusThreeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 19.1l3.9-3.875q.275-.275.688-.287t.712.287q.275.275.275.7t-.275.7L13.425 20.5q-.575.575-1.425.575t-1.425-.575L6.7 16.625q-.275-.275-.288-.688t.288-.712q.275-.275.7-.275t.7.275L12 19.1Zm0-5.95l3.9-3.875q.275-.275.688-.287t.712.287q.275.275.275.7t-.275.7l-3.875 3.875q-.575.575-1.425.575t-1.425-.575L6.7 10.675q-.275-.275-.288-.688t.288-.712Q6.975 9 7.4 9t.7.275L12 13.15Zm0-5.95l3.9-3.875q.275-.275.688-.287t.712.287q.275.275.275.7t-.275.7L13.425 8.6q-.575.575-1.425.575T10.575 8.6L6.7 4.725q-.275-.275-.288-.687t.288-.713q.275-.275.7-.275t.7.275L12 7.2Z"/>`),
		g.Group(children),
	)
}