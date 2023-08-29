package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.09 20.66a8.68 8.68 0 0 1-1.09-.07A8.8 8.8 0 0 1 3.41 13a8.71 8.71 0 0 1 6.83-9.67a1.23 1.23 0 0 1 1.27.48a1.27 1.27 0 0 1 .05 1.4a5.3 5.3 0 0 0-.66 3.47a5.24 5.24 0 0 0 4.38 4.38a5.19 5.19 0 0 0 3.47-.67a1.27 1.27 0 0 1 1.4.07a1.21 1.21 0 0 1 .48 1.26a8.7 8.7 0 0 1-8.54 6.94ZM10 5a7.25 7.25 0 1 0 9 9a6.74 6.74 0 0 1-9.61-5A6.75 6.75 0 0 1 10 5Z"/>`),
		g.Group(children),
	)
}