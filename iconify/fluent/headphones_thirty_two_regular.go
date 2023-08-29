package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.5 4C9.149 4 4 9.149 4 15.5V18h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H6.5A4.5 4.5 0 0 1 2 25.5v-10C2 8.044 8.044 2 15.5 2S29 8.044 29 15.5v10a4.5 4.5 0 0 1-4.5 4.5H20a1 1 0 0 1-1-1V19a1 1 0 0 1 1-1h7v-2.5C27 9.149 21.851 4 15.5 4ZM27 20h-6v8h3.5a2.5 2.5 0 0 0 2.5-2.5V20ZM4 20v5.5A2.5 2.5 0 0 0 6.5 28H10v-8H4Z"/>`),
		g.Group(children),
	)
}