package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmReelFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 216h-40.64A103.95 103.95 0 1 0 128 232h96a8 8 0 0 0 0-16ZM80 148a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm48 48a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm0-96a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm28 28a20 20 0 1 1 20 20a20 20 0 0 1-20-20Z"/>`),
		g.Group(children),
	)
}