package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeStand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 176a80 80 0 1 0-80-80a80.09 80.09 0 0 0 80 80Zm0-144a64 64 0 1 1-64 64a64.07 64.07 0 0 1 64-64Zm77.77 133.5a8 8 0 0 1-.23 11.32a111.24 111.24 0 0 1-69.54 30.9V224h24a8 8 0 0 1 0 16H96a8 8 0 0 1 0-16h24v-16.29A112 112 0 0 1 47.18 18.46a8 8 0 1 1 11.54 11.08a96 96 0 0 0 135.74 135.74a8 8 0 0 1 11.31.22Z"/>`),
		g.Group(children),
	)
}