package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 9V6a3 3 0 1 0-3 3h3zm0 0V7a2 2 0 1 1 2 2h-2zm-7 4v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-7m-7-3v11m8-8v-3a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v3h16z"/>`),
		g.Group(children),
	)
}