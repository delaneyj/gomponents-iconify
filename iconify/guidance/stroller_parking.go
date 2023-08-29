package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrollerParking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3 18.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm0 0v-2l1.693-1.395M23.5 6V5A2.5 2.5 0 0 0 21 2.5h-1L4.693 15.105M16 5.75L13.539 15.8a2.886 2.886 0 0 1-4.185 1.848l-4.66-2.543M0 14c.197 0 .384-.005.561-.014c1.642-.08 2.953-1.17 3.939-2.486h.25l2.608 1.422M17.75 20.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM10.85 7s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C12.746 6 11.15 7 11.15 7h-.3Z"/>`),
		g.Group(children),
	)
}