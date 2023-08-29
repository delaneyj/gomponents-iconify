package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeCoolOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.175 24L13 17.825V22h-2v-4.15l-3.25 3.2l-1.4-1.4l4.25-4.25l-2-2l-4.25 4.25l-1.4-1.4L6.15 13H2v-2h4.175l-5.4-5.4L2.2 4.175l18.4 18.4L19.175 24Zm.475-6.35L15 13h-1.125L11 10.125V9L6.35 4.35l1.4-1.4L11 6.15V2h2v4.15l3.25-3.2l1.4 1.4L13 9v2h2l4.65-4.65l1.4 1.4l-3.2 3.25H22v2h-4.15l3.2 3.25l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}