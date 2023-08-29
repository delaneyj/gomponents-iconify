package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTeardropFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 124a100.11 100.11 0 0 1-100 100H47.67A15.69 15.69 0 0 1 32 208.33V124a100 100 0 0 1 200 0Z"/>`),
		g.Group(children),
	)
}