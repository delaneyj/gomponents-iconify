package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornersOutBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 48v40a12 12 0 0 1-24 0V60h-28a12 12 0 0 1 0-24h40a12 12 0 0 1 12 12ZM88 196H60v-28a12 12 0 0 0-24 0v40a12 12 0 0 0 12 12h40a12 12 0 0 0 0-24Zm120-40a12 12 0 0 0-12 12v28h-28a12 12 0 0 0 0 24h40a12 12 0 0 0 12-12v-40a12 12 0 0 0-12-12ZM88 36H48a12 12 0 0 0-12 12v40a12 12 0 0 0 24 0V60h28a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}