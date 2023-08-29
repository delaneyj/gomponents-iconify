package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M88.04 30.319L75.124 17.401a2.42 2.42 0 0 0-3.419 0L37.392 51.714l-9.094-9.093a2.418 2.418 0 0 0-3.419 0L11.96 55.539a2.419 2.419 0 0 0 0 3.419L35.607 82.6a2.416 2.416 0 0 0 1.709.708c.029 0 .055-.016.083-.016c.024 0 .05.014.075.014c.621 0 1.236-.236 1.709-.708l48.857-48.86a2.416 2.416 0 0 0 0-3.419z"/>`),
		g.Group(children),
	)
}