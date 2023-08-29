package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectionInverseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 216a8 8 0 0 1-8 8h-32a8 8 0 0 1 0-16h32a8 8 0 0 1 8 8ZM40 152a8 8 0 0 0 8-8v-32a8 8 0 0 0-16 0v32a8 8 0 0 0 8 8Zm32 56H48v-24a8 8 0 0 0-16 0v24a16 16 0 0 0 16 16h24a8 8 0 0 0 0-16ZM224 48a16 16 0 0 0-16-16H48a15.87 15.87 0 0 0-10.66 4.11a7.67 7.67 0 0 0-1.23 1.23A15.87 15.87 0 0 0 32 48v24a8 8 0 0 0 16 0V59.31L196.69 208H184a8 8 0 0 0 0 16h24a15.91 15.91 0 0 0 10.66-4.1a7.35 7.35 0 0 0 .65-.59a6 6 0 0 0 .58-.65A15.87 15.87 0 0 0 224 208Z"/>`),
		g.Group(children),
	)
}