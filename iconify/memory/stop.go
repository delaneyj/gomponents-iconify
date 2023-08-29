package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 5v1h1v10h-1v1H6v-1H5V6h1V5h10M7 7v8h8V7H7Z"/>`),
		g.Group(children),
	)
}