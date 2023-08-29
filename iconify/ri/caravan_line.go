package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaravanLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.172 3a2 2 0 0 1 1.414.586l4.828 4.828A2 2 0 0 1 21 9.828V17h2v2h-8.126a4.002 4.002 0 0 1-7.748 0H3a1 1 0 0 1-1-1V5a2 2 0 0 1 2-2h10.172ZM11 16a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm3.172-11H4v12h3.126a4.002 4.002 0 0 1 7.748 0H19V9.828L14.172 5ZM14 7v6H6V7h8Zm-2 2H8v2h4V9Z"/>`),
		g.Group(children),
	)
}