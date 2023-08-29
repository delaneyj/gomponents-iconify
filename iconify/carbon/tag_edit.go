package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.707 7.293l-3-3a1 1 0 0 0-1.414 0L16 13.586V18h4.414l9.293-9.293a1 1 0 0 0 0-1.414zM19.586 16H18v-1.586l5-5L24.586 11zM26 9.586L24.414 8L26 6.414L27.586 8zM10 14a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 1.998 2.004A2.002 2.002 0 0 0 10 8z"/><path fill="currentColor" d="m27.006 14.235l-1.414 1.414L28 18.058L18.058 28L4 13.941V4h9.942l4.406 4.407l1.415-1.415l-4.407-4.406A2 2 0 0 0 13.94 2H4a2 2 0 0 0-2 2v9.941a2 2 0 0 0 .586 1.414l14.058 14.06a2.001 2.001 0 0 0 2.828 0l9.943-9.943a2.001 2.001 0 0 0 0-2.828Z"/>`),
		g.Group(children),
	)
}