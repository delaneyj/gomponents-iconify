package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkCircleSquareTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 7A5 5 0 1 0 2 7a5 5 0 0 0 10 0Zm1 0A6 6 0 1 1 1 7a6 6 0 0 1 12 0Zm-5 9v-2.07a6.953 6.953 0 0 0 1-.22V16a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2h-2.29a7.03 7.03 0 0 0 .22-1H16a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-5a3 3 0 0 1-3-3ZM9.854 5.146a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 1 1 .708-.708L6.5 7.793l2.646-2.647a.5.5 0 0 1 .708 0Zm3.5 10.208l3-3a.5.5 0 0 0-.708-.708L13 14.293l-1.146-1.147a.5.5 0 0 0-.708.708l1.5 1.5a.5.5 0 0 0 .708 0Z"/>`),
		g.Group(children),
	)
}