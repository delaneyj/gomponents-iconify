package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FramerLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 96V32a12 12 0 0 0-12-12H56a12 12 0 0 0-8 21l48.44 43H56a12 12 0 0 0-12 12v64a12 12 0 0 0 3.52 8.49l72 72A12 12 0 0 0 140 232v-60h60a12 12 0 0 0 8-21l-48.41-43H200a12 12 0 0 0 12-12Zm-43.56 52H128a12 12 0 0 0-12 12v43l-48-48v-47h55.44ZM188 84h-55.44l-45-40H188Z"/>`),
		g.Group(children),
	)
}