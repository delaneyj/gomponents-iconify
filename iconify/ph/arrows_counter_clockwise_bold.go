package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsCounterClockwiseBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 108H40a12 12 0 0 1-12-12V48a12 12 0 0 1 24 0v13.86C68.65 45.46 94.29 28 128 28c44.46 0 71.36 26.39 72.49 27.51a12 12 0 0 1-16.94 17C183.11 72.09 162 52 128 52c-29 0-50.94 17.68-64 32h24a12 12 0 0 1 0 24Zm128 40h-48a12 12 0 0 0 0 24h24c-13 14.32-35 32-64 32c-34 0-55.11-20.09-55.55-20.52a12 12 0 0 0-16.94 17C56.64 201.61 83.54 228 128 228c33.71 0 59.35-17.46 76-33.86V208a12 12 0 0 0 24 0v-48a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}