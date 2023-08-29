package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 11.414L28.586 10L24 14.586L19.414 10L18 11.414L22.586 16L18 20.585L19.415 22L24 17.414L28.587 22L30 20.587L25.414 16L30 11.414z"/><path fill="currentColor" d="M4 4a2 2 0 0 0-2 2v3.17a2 2 0 0 0 .586 1.415L10 18v8a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-2h-2v2h-4v-8.83l-.586-.585L4 9.171V6h20v2h2V6a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}