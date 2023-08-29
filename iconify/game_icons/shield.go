package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16c25 24 100 72 150 72v96c0 96-75 240-150 312c-75-72-150-216-150-312V88c50 0 125-48 150-72z"/>`),
		g.Group(children),
	)
}