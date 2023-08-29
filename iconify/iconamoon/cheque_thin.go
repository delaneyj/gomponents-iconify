package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChequeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5Zm4 9h3m-3-4h10"/>`),
		g.Group(children),
	)
}