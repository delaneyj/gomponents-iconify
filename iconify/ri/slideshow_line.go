package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideshowLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21v2h-2v-2H3a1 1 0 0 1-1-1V6h20v14a1 1 0 0 1-1 1h-8Zm-9-2h16V8H4v11Zm9-9h5v2h-5v-2Zm0 4h5v2h-5v-2Zm-4-4v3h3a3 3 0 1 1-3-3ZM2 3h20v2H2V3Z"/>`),
		g.Group(children),
	)
}