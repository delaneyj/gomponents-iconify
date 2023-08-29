package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GifBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 72v112a12 12 0 0 1-24 0V72a12 12 0 0 1 24 0Zm76 12a12 12 0 0 0 0-24h-48a12 12 0 0 0-12 12v112a12 12 0 0 0 24 0v-44h28a12 12 0 0 0 0-24h-28V84ZM96 116H72a12 12 0 0 0 0 24h12v12a20 20 0 0 1-40 0v-48a20 20 0 0 1 20-20a21.8 21.8 0 0 1 18.35 10.22a12 12 0 0 0 20.53-12.44A45.67 45.67 0 0 0 64 60a44.05 44.05 0 0 0-44 44v48a44 44 0 0 0 88 0v-24a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}