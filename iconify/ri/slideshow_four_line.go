package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideshowFourLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.17 3A3.001 3.001 0 0 1 11 1h2c1.306 0 2.418.835 2.83 2H21a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h5.17ZM4 5v14h16V5h-4.17A3.001 3.001 0 0 1 13 7h-2a3.001 3.001 0 0 1-2.83-2H4Zm7-2a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2Zm-1 6l5 3l-5 3V9Z"/>`),
		g.Group(children),
	)
}