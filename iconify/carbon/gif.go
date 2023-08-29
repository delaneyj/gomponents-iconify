package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 12v8a3 3 0 0 0 3 3h5v-8H6v2h2v4H5a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h5V9H5a3 3 0 0 0-3 3zm28-1V9h-8v14h2v-6h5v-2h-5v-4h6zM12 9v2h3v10h-3v2h8v-2h-3V11h3V9h-8z"/>`),
		g.Group(children),
	)
}