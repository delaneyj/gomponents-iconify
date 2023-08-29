package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leaderboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V9h5.5v12H2Zm7.25 0V3h5.5v18h-5.5Zm7.25 0V11H22v10h-5.5Z"/>`),
		g.Group(children),
	)
}