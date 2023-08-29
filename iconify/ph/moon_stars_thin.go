package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonStarsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 96a4 4 0 0 1-4 4h-20v20a4 4 0 0 1-8 0v-20h-20a4 4 0 0 1 0-8h20V72a4 4 0 0 1 8 0v20h20a4 4 0 0 1 4 4Zm-92-44h12v12a4 4 0 0 0 8 0V52h12a4 4 0 0 0 0-8h-12V32a4 4 0 0 0-8 0v12h-12a4 4 0 0 0 0 8Zm69.73 103.58a4 4 0 0 1 .71 4a92 92 0 1 1-118-118a4 4 0 0 1 5.29 4.54A93.18 93.18 0 0 0 100 64a92.1 92.1 0 0 0 92 92a93.18 93.18 0 0 0 17.91-1.74a4 4 0 0 1 3.82 1.32Zm-9.46 7.67A100 100 0 0 1 92.75 51.73a84 84 0 1 0 111.52 111.52Z"/>`),
		g.Group(children),
	)
}