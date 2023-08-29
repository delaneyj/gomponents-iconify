package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeStandLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 174a78 78 0 1 0-78-78a78.09 78.09 0 0 0 78 78Zm0-144a66 66 0 1 1-66 66a66.08 66.08 0 0 1 66-66Zm76.33 136.89a6 6 0 0 1-.17 8.48A109.21 109.21 0 0 1 134 205.83V226h26a6 6 0 0 1 0 12H96a6 6 0 0 1 0-12h26v-20.17A110 110 0 0 1 18 97.06a109.36 109.36 0 0 1 30.62-77.22a6 6 0 0 1 8.66 8.32a98 98 0 0 0 138.56 138.56a6 6 0 0 1 8.49.17Z"/>`),
		g.Group(children),
	)
}