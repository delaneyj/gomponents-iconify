package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkCircleSquareSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 10a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm1.854-5.146l-2 2a.5.5 0 0 1-.708 0l-1-1a.5.5 0 1 1 .708-.708L5 5.793l1.646-1.647a.5.5 0 1 1 .708.708ZM6 10.978V12.5A2.5 2.5 0 0 0 8.5 15h4a2.5 2.5 0 0 0 2.5-2.5v-4A2.5 2.5 0 0 0 12.5 6h-1.522a5.48 5.48 0 0 1-.185 1H12.5A1.5 1.5 0 0 1 14 8.5v4a1.5 1.5 0 0 1-1.5 1.5h-4A1.5 1.5 0 0 1 7 12.5v-1.707a5.488 5.488 0 0 1-1 .185Zm6.353-1.124a.5.5 0 0 0-.708-.708L10 10.793l-.646-.647a.5.5 0 0 0-.708.708l1 1a.5.5 0 0 0 .708 0l2-2Z"/>`),
		g.Group(children),
	)
}