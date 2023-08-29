package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListChecksBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 128a12 12 0 0 1-12 12h-88a12 12 0 0 1 0-24h88a12 12 0 0 1 12 12ZM128 76h88a12 12 0 0 0 0-24h-88a12 12 0 0 0 0 24Zm88 104h-88a12 12 0 0 0 0 24h88a12 12 0 0 0 0-24ZM79.51 39.51L56 63l-7.51-7.52a12 12 0 0 0-17 17l16 16a12 12 0 0 0 17 0l32-32a12 12 0 0 0-17-17Zm0 64L56 127l-7.51-7.52a12 12 0 1 0-17 17l16 16a12 12 0 0 0 17 0l32-32a12 12 0 0 0-17-17Zm0 64L56 191l-7.51-7.52a12 12 0 1 0-17 17l16 16a12 12 0 0 0 17 0l32-32a12 12 0 0 0-17-17Z"/>`),
		g.Group(children),
	)
}