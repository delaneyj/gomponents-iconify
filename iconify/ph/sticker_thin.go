package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickerThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 36H88a52.06 52.06 0 0 0-52 52v80a52.06 52.06 0 0 0 52 52h48a3.83 3.83 0 0 0 1.26-.21c25.46-8.48 74.05-57.07 82.53-82.53A3.83 3.83 0 0 0 220 136V88a52.06 52.06 0 0 0-52-52ZM44 168V88a44.05 44.05 0 0 1 44-44h80a44.05 44.05 0 0 1 44 44v44h-28a52.06 52.06 0 0 0-52 52v28H88a44.05 44.05 0 0 1-44-44Zm96 42v-26a44.05 44.05 0 0 1 44-44h26c-11.1 22.23-47.77 58.9-70 70Z"/>`),
		g.Group(children),
	)
}