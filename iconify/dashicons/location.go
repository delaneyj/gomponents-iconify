package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Location(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2C6.69 2 4 4.69 4 8c0 2.02 1.17 3.71 2.53 4.89c.43.37 1.18.96 1.85 1.83c.74.97 1.41 2.01 1.62 2.71c.21-.7.88-1.74 1.62-2.71c.67-.87 1.42-1.46 1.85-1.83C14.83 11.71 16 10.02 16 8c0-3.31-2.69-6-6-6zm0 2.56a3.44 3.44 0 1 1 0 6.88a3.44 3.44 0 0 1 0-6.88z"/>`),
		g.Group(children),
	)
}