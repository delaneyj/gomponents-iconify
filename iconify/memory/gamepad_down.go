package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 14H9v4h4v-4Z"/>`),
		g.Group(children),
	)
}