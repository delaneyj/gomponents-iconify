package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 21V9h-2v14h8v-2h-6zM18 9l-1.52 5l-.48 1.98l-.46-1.98L14 9h-2v14h2v-8l-.16-2l.58 2L16 19.63L17.58 15l.58-2l-.16 2v8h2V9h-2zm-8 0H8l-2 6l-2-6H2l2.75 7L2 23h2l2-6l2 6h2l-2.75-7L10 9z"/>`),
		g.Group(children),
	)
}