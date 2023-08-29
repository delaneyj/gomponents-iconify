package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 7H9v1H8v4h1v1h4v-1h1V9h-1V8Z"/>`),
		g.Group(children),
	)
}