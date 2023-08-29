package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FogBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.762 7.647C6.762 4.528 9.32 2 12.476 2c2.94 0 5.361 2.194 5.68 5.015C20.392 7.78 22 9.881 22 12.353a5.57 5.57 0 0 1-.808 2.897H22a.75.75 0 0 1 0 1.5H2a.75.75 0 0 1 0-1.5h.271A4.182 4.182 0 0 1 2 13.765c0-2.34 1.919-4.236 4.286-4.236c.284 0 .562.028.83.08a5.576 5.576 0 0 1-.354-1.962ZM5 18.25a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 0-1.5H5Zm3 3a.75.75 0 0 0 0 1.5h8a.75.75 0 0 0 0-1.5H8Z"/>`),
		g.Group(children),
	)
}