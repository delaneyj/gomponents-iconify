package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpFileUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 10h4v6h6v-6h4l-7-7l-7 7zm0 8v2h14v-2H5z"/>`),
		g.Group(children),
	)
}