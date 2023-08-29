package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.47 2.18a6.49 6.49 0 0 1 3.22 5.5a6.84 6.84 0 0 1-7.08 6.55a7.42 7.42 0 0 1-3.82-1a7 7 0 0 1-1.9-1.65a9.47 9.47 0 0 0 3.34-.74a9.92 9.92 0 0 0 3.3-2.24a10 10 0 0 0 2.94-6.42M10.82.45a.66.66 0 0 0-.66.69a8.63 8.63 0 0 1-2.55 6.54a8.68 8.68 0 0 1-6.09 2.57H.66a.66.66 0 0 0-.58 1a8.45 8.45 0 0 0 7.53 4.39A8.15 8.15 0 0 0 16 7.68A7.85 7.85 0 0 0 11.07.51a.61.61 0 0 0-.25-.06z"/>`),
		g.Group(children),
	)
}