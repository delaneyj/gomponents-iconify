package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbstractFiftyEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21 21v228.569h168.172v115.56h-21.003V270.98H21V491h470V270.98H343.831v94.147h-21.003v-115.56H491V21h-35.103v145.113H347.063V21H310.49v145.113H201.656V21h-36.719v145.113H56.103V21z"/>`),
		g.Group(children),
	)
}