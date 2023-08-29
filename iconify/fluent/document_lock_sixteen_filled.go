package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLockSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 1v3.5A1.5 1.5 0 0 0 10.5 6H14v7.5a1.5 1.5 0 0 1-1.5 1.5H9.732A1.99 1.99 0 0 0 10 14v-4a2 2 0 0 0-1.5-1.937V8A3 3 0 0 0 4 5.401V2.5A1.5 1.5 0 0 1 5.5 1H9Zm1 .25V4.5a.5.5 0 0 0 .5.5h3.25L10 1.25ZM3.5 8v1H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-.5V8a2 2 0 1 0-4 0Zm1 1V8a1 1 0 0 1 2 0v1h-2Zm1 2.25a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}