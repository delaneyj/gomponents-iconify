package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserFavoriteAltFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.494 3a3.474 3.474 0 0 0-2.48 1.04l-.511.522l-.516-.523a3.48 3.48 0 0 0-4.96 0a3.59 3.59 0 0 0 0 5.025l5.476 5.543l5.472-5.543a3.59 3.59 0 0 0 0-5.025A3.474 3.474 0 0 0 26.494 3zM16 30h-2v-5a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v5H2v-5a5.006 5.006 0 0 1 5-5h4a5.006 5.006 0 0 1 5 5zM9 10a3 3 0 1 1-3 3a3 3 0 0 1 3-3m0-2a5 5 0 1 0 5 5a5 5 0 0 0-5-5z"/>`),
		g.Group(children),
	)
}