package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkCircleSquareTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 13A6 6 0 1 0 7 1a6 6 0 0 0 0 12Zm2.854-7.146l-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 1 1 .708-.708L6.5 7.793l2.646-2.647a.5.5 0 1 1 .708.708ZM8 13.929V16a3 3 0 0 0 3 3h5a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3h-2.07a6.953 6.953 0 0 1-.22 1H16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-5a2 2 0 0 1-2-2v-2.29a7.03 7.03 0 0 1-1 .22Zm8.354-1.575a.5.5 0 0 0-.708-.708L13 14.293l-1.146-1.147a.5.5 0 0 0-.708.708l1.5 1.5a.5.5 0 0 0 .708 0l3-3Z"/>`),
		g.Group(children),
	)
}