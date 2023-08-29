package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BathtubF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.487 2.897A3 3 0 0 1 11 5.857v2H5v-2c0-1.038.528-1.954 1.33-2.492A2 2 0 0 0 3 4.858v6h16a1 1 0 0 1 0 2v1a6.002 6.002 0 0 1-4 5.658V20a1 1 0 1 1-2 0v-.142H7V20a1 1 0 1 1-2 0v-.484a6.002 6.002 0 0 1-4-5.658v-1a1 1 0 0 1 0-2v-6a4 4 0 0 1 7.487-1.96z"/>`),
		g.Group(children),
	)
}