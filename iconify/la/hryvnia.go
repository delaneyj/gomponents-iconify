package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hryvnia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.21 5c-.99 0-1.948.29-2.773.84L9.445 7.168l1.11 1.664l1.992-1.328A2.99 2.99 0 0 1 14.21 7h4.385a2.407 2.407 0 0 1 1.539 4.254L19.24 12H8v2h8.84l-3.602 3H8v2h2.84l-.254.21A4.398 4.398 0 0 0 9 22.597A4.409 4.409 0 0 0 13.404 27h4.385c.99 0 1.95-.29 2.773-.84l1.993-1.328l-1.11-1.664l-1.992 1.328A2.99 2.99 0 0 1 17.79 25h-4.385a2.407 2.407 0 0 1-1.539-4.254L13.961 19H24v-2h-7.639l3.6-3H24v-2h-1.85A4.39 4.39 0 0 0 23 9.404A4.409 4.409 0 0 0 18.596 5H14.21z"/>`),
		g.Group(children),
	)
}