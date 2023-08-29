package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flashy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.26 22.32a1.71 1.71 0 1 0 2.42 2.42"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.28 13.26l-.61 3.49a4.38 4.38 0 0 1-.19.82a2 2 0 0 1-.26.5a4.39 4.39 0 0 1-.57.62L5.79 36.55a1 1 0 0 0 0 1.42L10 42.21a1 1 0 0 0 1.42 0l17.89-17.86a4.39 4.39 0 0 1 .62-.57a2 2 0 0 1 .5-.26a4.38 4.38 0 0 1 .82-.19l3.49-.61A10.78 10.78 0 0 0 39.31 20L28 8.69a10.78 10.78 0 0 0-2.72 4.57Zm-2 9.06a1.71 1.71 0 0 1 2.42 2.42l-2.42 2.43a1.72 1.72 0 0 1-2.43-2.43ZM39.31 20a43 43 0 0 0 3-3.27a2.28 2.28 0 0 0 0-1.74A5.7 5.7 0 0 0 41 13.44L34.56 7A5.7 5.7 0 0 0 33 5.67a2.28 2.28 0 0 0-1.74 0a43 43 0 0 0-3.27 3"/>`),
		g.Group(children),
	)
}