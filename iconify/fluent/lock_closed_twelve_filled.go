package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosedTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 4V3a2 2 0 1 0-4 0v1h-.5A1.5 1.5 0 0 0 2 5.5v4A1.5 1.5 0 0 0 3.5 11h5A1.5 1.5 0 0 0 10 9.5v-4A1.5 1.5 0 0 0 8.5 4H8ZM6 2a1 1 0 0 1 1 1v1H5V3a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}