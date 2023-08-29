package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfinityLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M246 128a54 54 0 0 1-92.18 38.18a3.07 3.07 0 0 1-.25-.26l-60-67.74a42 42 0 1 0 0 59.64l8.57-9.67a6 6 0 1 1 9 8l-8.69 9.81a3.07 3.07 0 0 1-.25.26a54 54 0 1 1 0-76.36a3.07 3.07 0 0 1 .25.26l60 67.74a42 42 0 1 0 0-59.64l-8.57 9.67a6 6 0 1 1-9-8l8.69-9.81a3.07 3.07 0 0 1 .25-.26A54 54 0 0 1 246 128Z"/>`),
		g.Group(children),
	)
}