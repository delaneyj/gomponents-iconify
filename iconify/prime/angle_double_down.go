package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.25a.74.74 0 0 1-.53-.22L8 13.53a.75.75 0 0 1 1-1.06l3 3l3-3a.75.75 0 0 1 1 1.06L12.53 17a.74.74 0 0 1-.53.25Zm0-5.5a.74.74 0 0 1-.53-.22L8 8a.75.75 0 0 1 1-1l3 3l3-3a.75.75 0 0 1 1 1l-3.5 3.5a.74.74 0 0 1-.5.25Z"/>`),
		g.Group(children),
	)
}