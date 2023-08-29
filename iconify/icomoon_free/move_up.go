package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 8v6h1V8h2.5l-3-3l-3 3zM1 3h1.5v1H1V3zm2 0h1.5v1H3V3zm2 0h1v1.5H5V3zM1 6.5h1V8H1V6.5zm1.5.5H4v1H2.5V7zm2 0H6v1H4.5V7zM1 4.5h1V6H1V4.5zM5 5h1v1.5H5V5zm0 6v3H2v-3h3zm1-1H1v5h5v-5z"/>`),
		g.Group(children),
	)
}