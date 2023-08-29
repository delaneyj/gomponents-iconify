package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.47 2.22a.75.75 0 0 1 1.06 0c.523.523 2.494 2.614 4.34 5.316c1.823 2.669 3.63 6.082 3.63 9.214c0 3.041-.917 5.374-2.49 6.947C18.442 25.267 16.299 26 14 26c-2.298 0-4.441-.733-6.01-2.303C6.417 22.124 5.5 19.79 5.5 16.75c0-3.132 1.807-6.545 3.63-9.214c1.846-2.702 3.817-4.793 4.34-5.316Z"/>`),
		g.Group(children),
	)
}