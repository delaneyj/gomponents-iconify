package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmReelBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 212h-28.2a107.94 107.94 0 1 0-67.8 24h96a12 12 0 0 0 0-24ZM44 128a84 84 0 1 1 84 84a84.09 84.09 0 0 1-84-84Zm64-44a20 20 0 1 1 20 20a20 20 0 0 1-20-20Zm20 108a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm24-64a20 20 0 1 1 20 20a20 20 0 0 1-20-20Zm-48 0a20 20 0 1 1-20-20a20 20 0 0 1 20 20Z"/>`),
		g.Group(children),
	)
}