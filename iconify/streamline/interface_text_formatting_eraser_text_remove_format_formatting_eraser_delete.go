package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingEraserTextRemoveFormatFormattingEraserDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 13.5h10m-1.18-7.29a.94.94 0 0 0 0-1.35L8.25.78a1 1 0 0 0-1.36 0L.78 6.89a1 1 0 0 0 0 1.36l2.44 2.43a.92.92 0 0 0 .67.29h3.28a.92.92 0 0 0 .68-.29Z"/>`),
		g.Group(children),
	)
}