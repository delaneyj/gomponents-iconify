package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 106a53.74 53.74 0 0 0-19.94 3.83L141.23 51a6 6 0 1 0-10.46-5.89l-49.54 88A54 54 0 1 0 128 106Zm0 96a42 42 0 1 1 42-42a42 42 0 0 1-42 42Z"/>`),
		g.Group(children),
	)
}