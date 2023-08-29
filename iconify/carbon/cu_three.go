package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CuThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M10 23H5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h5v2H5v6h5z" fill="currentColor"/><path d="M18 23h-4a2 2 0 0 1-2-2V9h2v12h4V9h2v12a2 2 0 0 1-2 2z" fill="currentColor"/><path d="M28 9h-6v2h6v4h-5v2h5v4h-6v2h6a2 2 0 0 0 2-2V11a2 2 0 0 0-2-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}