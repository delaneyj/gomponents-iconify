package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineCompress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19h3v3h2v-3h3l-4-4l-4 4zm8-15h-3V1h-2v3H8l4 4l4-4zM4 9v2h16V9H4zm0 3h16v2H4z"/>`),
		g.Group(children),
	)
}