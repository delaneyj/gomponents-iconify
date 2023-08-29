package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarcodeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 40H32a8 8 0 0 0-8 8v160a8 8 0 0 0 8 8h192a8 8 0 0 0 8-8V48a8 8 0 0 0-8-8ZM40 64a8 8 0 0 1 8-8h32a8 8 0 0 1 0 16H56v24a8 8 0 0 1-16 0Zm40 136H48a8 8 0 0 1-8-8v-32a8 8 0 0 1 16 0v24h24a8 8 0 0 1 0 16Zm24-48a8 8 0 0 1-16 0v-48a8 8 0 0 1 16 0Zm32 0a8 8 0 0 1-16 0v-48a8 8 0 0 1 16 0Zm24 8a8 8 0 0 1-8-8v-48a8 8 0 0 1 16 0v48a8 8 0 0 1-8 8Zm56 32a8 8 0 0 1-8 8h-32a8 8 0 0 1 0-16h24v-24a8 8 0 0 1 16 0Zm0-96a8 8 0 0 1-16 0V72h-24a8 8 0 0 1 0-16h32a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}