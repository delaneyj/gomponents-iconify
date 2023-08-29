package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceEqTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2a.5.5 0 0 1 .5.5v15a.5.5 0 0 1-1 0v-15A.5.5 0 0 1 10 2Zm3.5 3a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-1 0v-9a.5.5 0 0 1 .5-.5ZM7 5.5a.5.5 0 0 0-1 0v9a.5.5 0 0 0 1 0v-9ZM16.5 8a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 .5-.5ZM4 8.5a.5.5 0 0 0-1 0v3a.5.5 0 0 0 1 0v-3Z"/>`),
		g.Group(children),
	)
}