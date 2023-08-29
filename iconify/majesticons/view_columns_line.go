package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumnsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h4M9 6v12M9 6h6M9 18h6m0 0h4a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-4m0 12V6"/>`),
		g.Group(children),
	)
}