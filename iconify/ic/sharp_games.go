package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpGames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 7.5V2H9v5.5l3 3l3-3zM7.5 9H2v6h5.5l3-3l-3-3zM9 16.5V22h6v-5.5l-3-3l-3 3zM16.5 9l-3 3l3 3H22V9h-5.5z"/>`),
		g.Group(children),
	)
}