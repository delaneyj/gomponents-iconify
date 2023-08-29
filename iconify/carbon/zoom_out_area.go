package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 19h8v2h-8z"/><path fill="currentColor" d="m31 29.586l-4.689-4.688a8.028 8.028 0 1 0-1.414 1.414L29.586 31zM20 26a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6zM4 8H2V4a2.002 2.002 0 0 1 2-2h4v2H4zm22 0h-2V4h-4V2h4a2.002 2.002 0 0 1 2 2zM12 2h4v2h-4zM8 26H4a2.002 2.002 0 0 1-2-2v-4h2v4h4zM2 12h2v4H2z"/>`),
		g.Group(children),
	)
}