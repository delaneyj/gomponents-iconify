package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsBookTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22H6a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h14a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1Zm-1-2v-2H6a1 1 0 1 0 0 2h13Zm-7-10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-3 4h6a3 3 0 1 0-6 0Z"/>`),
		g.Group(children),
	)
}