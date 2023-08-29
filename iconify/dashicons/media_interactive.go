package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaInteractive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12 2l4 4v12H4V2h8zm0 4h3l-3-3v3zm2 8V8H6v6h3l-1 2h1l1-2l1 2h1l-1-2h3zm-6-3c-.55 0-1-.45-1-1s.45-1 1-1s1 .45 1 1s-.45 1-1 1zm5-2v2h-3V9h3zm0 3v1H7v-1h6z"/>`),
		g.Group(children),
	)
}