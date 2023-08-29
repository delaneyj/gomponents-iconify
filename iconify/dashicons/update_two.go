package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpdateTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M10 3c2.09 0 3.97.92 5.25 2.38l-1.71 1.95A4.44 4.44 0 0 0 10 5.55C7.89 5.55 6.13 7.03 5.68 9H8l-3.5 4L1 9h2.08C3.57 5.61 6.47 3 10 3zm0 14.04c-2.1 0-3.97-.92-5.25-2.38l1.71-1.95A4.44 4.44 0 0 0 10 14.49c2.11 0 3.87-1.48 4.32-3.45H12l3.5-4 3.5 4h-2.08c-.49 3.39-3.4 6-6.92 6z" fill="currentColor"/>`),
		g.Group(children),
	)
}