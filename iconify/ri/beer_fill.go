package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3a2 2 0 0 1 1.972 2.335l1.973.33a4.011 4.011 0 0 0-.005-1.361A2 2 0 0 1 15.733 7H5a1 1 0 1 1 .539-1.843a1 1 0 0 0 1.513-.614A2.001 2.001 0 0 1 9 3Zm1.516-1.703A3.998 3.998 0 0 0 5.51 3.043A3 3 0 0 0 3 8.236V20a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2h2a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2h-2v-.354a4 4 0 0 0-4.896-6.169a4.01 4.01 0 0 0-1.588-1.18ZM17 18v-7h2v7h-2ZM7 11h2v7H7v-7Zm4 0h2v7h-2v-7Z"/>`),
		g.Group(children),
	)
}