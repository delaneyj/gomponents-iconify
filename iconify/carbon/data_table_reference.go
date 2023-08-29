package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTableReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm16 2h4v2h-4zm-4-4h4v2h-4zm4-4h4v2h-4zM8 14h4v2H8z"/><path fill="currentColor" d="M27 3H5a2.003 2.003 0 0 0-2 2v11h2v-5h22v16H16v2h11a2.003 2.003 0 0 0 2-2V5a2.003 2.003 0 0 0-2-2Zm0 6H5V5h22Z"/>`),
		g.Group(children),
	)
}