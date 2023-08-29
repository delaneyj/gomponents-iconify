package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsBottleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.435 2.06H6.565a2.5 2.5 0 0 0-2.5 2.5v2a1.492 1.492 0 0 0 1.22 1.47v11.41a2.5 2.5 0 0 0 2.5 2.5h8.43a2.5 2.5 0 0 0 2.5-2.5V8.03a1.492 1.492 0 0 0 1.22-1.47v-2a2.5 2.5 0 0 0-2.5-2.5Zm.28 17.38a1.5 1.5 0 0 1-1.5 1.5h-8.43a1.5 1.5 0 0 1-1.5-1.5v-.88h3.52a.491.491 0 0 0 .48-.5a.485.485 0 0 0-.48-.5h-3.52V15h2.57a.5.5 0 0 0 0-1h-2.57v-2.55h3.52a.491.491 0 0 0 .48-.5a.485.485 0 0 0-.48-.5h-3.52V8.06h11.43Zm1.22-12.88a.5.5 0 0 1-.5.5H5.565a.5.5 0 0 1-.5-.5v-2a1.5 1.5 0 0 1 1.5-1.5h10.87a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		g.Group(children),
	)
}