package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AncientPavilionLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.516 7.877A4.5 4.5 0 0 0 5 11.887v7.114H3v2h18v-2h-2v-7.114a4.5 4.5 0 0 0 3.484-4.01A9 9 0 0 1 12.514 2h-1.029a9 9 0 0 1-9.97 5.877ZM17 19H7v-7h10v7.001Zm1.555-9.064l-.274.063H5.72l-.275-.063a2.593 2.593 0 0 1-.39-.121C6.59 9.532 8.021 8.92 9.35 7.98A10.88 10.88 0 0 0 12 5.346a10.877 10.877 0 0 0 2.649 2.635c1.328.94 2.76 1.551 4.298 1.835c-.125.05-.256.09-.392.121Z"/>`),
		g.Group(children),
	)
}