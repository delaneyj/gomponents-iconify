package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h11a2 2 0 0 1 2 2v11m-1.172 2.821A1.993 1.993 0 0 1 19 20H5a2 2 0 0 1-2-2V6a2 2 0 0 1 1.156-1.814M7 8h1m4 0h5M7 12h5m4 0h1M7 16h9M3 3l18 18"/>`),
		g.Group(children),
	)
}