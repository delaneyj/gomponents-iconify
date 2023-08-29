package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wmv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29 9l-2 13l-2-13h-2l2.52 14h2.96L31 9h-2zM19 9l-1.52 5l-.48 1.98l-.46-1.98L15 9h-2v14h2v-8l-.16-2l.58 2L17 19.63L18.58 15l.58-2l-.16 2v8h2V9h-2zM9.2 9l-.34 8l-.26 4.54L8.19 18l-.68-5.46H5.49L4.81 18l-.41 3.54L4.14 17L3.8 9H2l1 14h2.27l.76-4.93l.46-4.07l.01-.03l.01.03l.46 4.07l.76 4.93H10l1-14H9.2z"/>`),
		g.Group(children),
	)
}