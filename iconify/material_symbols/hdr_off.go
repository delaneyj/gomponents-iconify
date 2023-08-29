package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L1.4 4.2l1.4-1.4l18.4 18.35l-1.4 1.45Zm-.3-7.6l-.9-2h-1.1v1.65l-1.5-1.5V9h3.5q.6 0 1.05.45T21 10.5v1q0 .45-.263.813t-.637.587L21 15h-1.5Zm-2-3.5h2v-1h-2v1Zm-3 .15L11.85 9H13q.6 0 1.05.45t.45 1.05v1.15ZM3 15V9h1.5v2h2V9H8v6H6.5v-2.5h-2V15H3Zm6.5-4.1l1.5 1.5v1.1h1.125l1.375 1.4q-.125.05-.25.075T13 15H9.5v-4.1Z"/>`),
		g.Group(children),
	)
}