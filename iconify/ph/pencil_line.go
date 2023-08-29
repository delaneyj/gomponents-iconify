package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m227.32 73.37l-44.69-44.68a16 16 0 0 0-22.63 0L36.69 152A15.86 15.86 0 0 0 32 163.31V208a16 16 0 0 0 16 16h168a8 8 0 0 0 0-16H115.32l112-112a16 16 0 0 0 0-22.63ZM136 75.31L152.69 92L68 176.69L51.31 160ZM48 208v-28.69L76.69 208Zm48-3.31L79.32 188L164 103.31L180.69 120Zm96-96L147.32 64l24-24L216 84.69Z"/>`),
		g.Group(children),
	)
}