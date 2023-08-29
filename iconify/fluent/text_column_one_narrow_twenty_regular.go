package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColumnOneNarrowTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 5a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5Zm0 3a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5ZM7 11.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5Zm.5 2.5a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5Z"/>`),
		g.Group(children),
	)
}