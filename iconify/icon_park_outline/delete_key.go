package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18.424 10.538A2 2 0 0 1 19.788 10H42a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2H19.788a2 2 0 0 1-1.364-.538L4 24l14.424-13.462ZM36 19L26 29m0-10l10 10"/>`),
		g.Group(children),
	)
}