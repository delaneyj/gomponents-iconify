package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectSquareF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4zm7.746 5.413a2.042 2.042 0 0 0-.575-.54c-.965-.606-2.27-.365-2.917.54L5.356 9.468A1.887 1.887 0 0 0 5 10.567c0 1.089.941 1.972 2.102 1.972h5.796c.417 0 .824-.116 1.17-.334c.965-.607 1.222-1.832.576-2.737l-2.898-4.055zM6 13.54a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2H6zm1.102-2.972L10 6.512l2.898 4.055H7.102z"/>`),
		g.Group(children),
	)
}