package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 16.25a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l3-3l-3-3A.75.75 0 0 1 13.53 8l3.5 3.5a.75.75 0 0 1 0 1.06L13.53 16a.74.74 0 0 1-.53.25Zm-5.5 0A.74.74 0 0 1 7 16a.75.75 0 0 1 0-1l3-3l-3-3a.75.75 0 0 1 1-1l3.5 3.5a.75.75 0 0 1 0 1.06L8 16a.74.74 0 0 1-.5.25Z"/>`),
		g.Group(children),
	)
}