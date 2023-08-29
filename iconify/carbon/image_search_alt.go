package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSearchAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 16a5.99 5.99 0 0 0 9.471 4.885L24.586 26L26 24.586l-5.115-5.115A5.997 5.997 0 1 0 10 16Zm2 0a4 4 0 1 1 4 4a4.005 4.005 0 0 1-4-4Z"/><path fill="currentColor" d="M29 7h-6.46l-1.71-2.55A.995.995 0 0 0 20 4h-8a.995.995 0 0 0-.83.45L9.46 7H3a1.003 1.003 0 0 0-1 1v17a1.003 1.003 0 0 0 1 1h9v-2H4V9h6a.995.995 0 0 0 .83-.45L12.54 6h6.92l1.71 2.55A.995.995 0 0 0 22 9h6v12h2V8a1.003 1.003 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}