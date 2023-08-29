package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M151.62 119.45a46 46 0 1 0-47.24 0a54 54 0 1 0 47.24 0ZM94 80a34 34 0 1 1 34 34a34 34 0 0 1-34-34Zm34 130a42 42 0 1 1 42-42a42 42 0 0 1-42 42Z"/>`),
		g.Group(children),
	)
}