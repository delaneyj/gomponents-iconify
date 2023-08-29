package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 42a54 54 0 1 0 19.94 104.17l-33.17 58.88a6 6 0 1 0 10.46 5.89l49.54-88A54 54 0 0 0 128 42Zm0 96a42 42 0 1 1 42-42a42 42 0 0 1-42 42Z"/>`),
		g.Group(children),
	)
}