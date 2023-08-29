package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAllTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Zm10 13h5a1 1 0 0 0 1-1v-5h-6v6Zm-2-6H5v5a1 1 0 0 0 1 1h5v-6Zm2-2h6V6a1 1 0 0 0-1-1h-5v6Zm-2-6H6a1 1 0 0 0-1 1v5h6V5Z"/>`),
		g.Group(children),
	)
}