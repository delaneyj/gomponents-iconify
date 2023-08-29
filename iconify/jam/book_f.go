package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18V1H3v17H0V0h11a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5zM6 4v2h5V4H6zm0 3v2h5V7H6zM3 18h2v2H3v-2z"/>`),
		g.Group(children),
	)
}