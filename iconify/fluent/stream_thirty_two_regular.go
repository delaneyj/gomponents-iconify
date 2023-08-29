package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreamThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.5 3C14.701 3 10 7.701 10 13.5v1A2.5 2.5 0 0 1 7.5 17H4a1 1 0 1 0 0 2h3.5a4.5 4.5 0 0 0 4.5-4.5v-1A8.5 8.5 0 0 1 20.5 5H23a1 1 0 1 0 0-2h-2.5Zm-6 11.5A6.5 6.5 0 0 1 21 8h7a1 1 0 1 1 0 2h-7a4.5 4.5 0 0 0-4.5 4.5v2A7.5 7.5 0 0 1 9 24H4a1 1 0 1 1 0-2h5a5.5 5.5 0 0 0 5.5-5.5v-2ZM9 27h1.163A8.837 8.837 0 0 0 19 18.163V16a3 3 0 0 1 3-3h6a1 1 0 1 1 0 2h-6a1 1 0 0 0-1 1v2.313c0 .021 0 .043-.002.065C20.883 24.264 16.076 29 10.163 29H9a1 1 0 1 1 0-2Z"/>`),
		g.Group(children),
	)
}