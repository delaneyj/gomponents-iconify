package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighResSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 18.5h4V17h-2.5v-.75h2v-1.5h-2V14H14v-1.5h-4v6Zm5 0h4.5V15h-3v-1h3v-1.5H15V16h3v1h-3v1.5Zm-7.5-7H9V9.25h1.5v2.25H12v-6h-1.5v2.25H9V5.5H7.5v6Zm7 0H16v-6h-1.5v6ZM1 21V3h22v18H1Zm3.5-2.5H6v-2h.6l.9 2H9l-.9-2.1q.375-.225.638-.588T9 15v-1q0-.625-.438-1.063T7.5 12.5h-3v6ZM6 15v-1h1.5v1H6Z"/>`),
		g.Group(children),
	)
}