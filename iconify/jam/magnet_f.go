package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 12c0-3.866-2.239-7-5-7s-5 3.134-5 7H0C0 5.373 4.477 0 10 0s10 5.373 10 12h-5zM0 14h5v5a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1v-5zm15 0h5v5a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1v-5z"/>`),
		g.Group(children),
	)
}