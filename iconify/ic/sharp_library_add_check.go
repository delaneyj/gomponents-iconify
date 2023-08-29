package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpLibraryAddCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2H6v16h16V2zm-9.53 12L9 10.5l1.4-1.41l2.07 2.08L17.6 6L19 7.41L12.47 14zM4 6H2v16h16v-2H4V6z"/>`),
		g.Group(children),
	)
}