package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4h-5V1h-2v3H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h5v3h2v-3h5a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zM6 18V6h5v12H6z"/>`),
		g.Group(children),
	)
}