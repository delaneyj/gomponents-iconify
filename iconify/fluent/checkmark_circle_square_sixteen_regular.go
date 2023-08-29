package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkCircleSquareSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 5.5a3.5 3.5 0 1 0-7 0a3.5 3.5 0 0 0 7 0Zm1 0a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-3.999 7v-1.522a5.488 5.488 0 0 0 1-.185V12.5a1.5 1.5 0 0 0 1.5 1.5h4a1.5 1.5 0 0 0 1.5-1.5v-4a1.5 1.5 0 0 0-1.5-1.5h-1.707a5.48 5.48 0 0 0 .185-1H12.5A2.5 2.5 0 0 1 15 8.5v4a2.5 2.5 0 0 1-2.5 2.5h-4A2.5 2.5 0 0 1 6 12.5Zm1.353-8.354a.5.5 0 0 1 0 .708l-2 2a.5.5 0 0 1-.708 0l-1-1a.5.5 0 1 1 .708-.708L5 5.793l1.646-1.647a.5.5 0 0 1 .708 0Zm5 5.708a.5.5 0 0 0-.708-.708L10 10.793l-.646-.647a.5.5 0 0 0-.708.708l1 1a.5.5 0 0 0 .708 0l2-2Z"/>`),
		g.Group(children),
	)
}