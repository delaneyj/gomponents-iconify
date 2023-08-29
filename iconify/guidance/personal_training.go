package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalTraining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M17.5 13.599a13.28 13.28 0 0 0-1-.099a3 3 0 0 0-3 3v1h4.677a1.323 1.323 0 0 0 1.276-1.67L18 10.436M7 20.5h17m-17 3c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1l-4-3.25v-.25l1.789-5.009A3 3 0 0 1 5.114 6.5H6.5v3.446a1.554 1.554 0 0 0 2.378 1.318L12.5 9M.5 23.75S2 21 2 17.5m21.85-3.118a399.79 399.79 0 0 1-3.612-.439l3.612.439ZM11.5 15.9s-1 1.6-2.25 1.6a1.75 1.75 0 0 1-1.75-1.75c0-.966.783-1.746 1.75-1.746c1.25 0 2.25 1.596 2.25 1.596v.3Zm6-5.4a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-10.4-6s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C8.996 3.5 7.4 4.5 7.4 4.5h-.3Z"/>`),
		g.Group(children),
	)
}