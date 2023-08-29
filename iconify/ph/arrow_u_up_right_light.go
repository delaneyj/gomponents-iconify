package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUUpRightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M171.76 131.76L209.51 94H88a50 50 0 0 0 0 100h88a6 6 0 0 1 0 12H88a62 62 0 0 1 0-124h121.51l-37.75-37.76a6 6 0 0 1 8.48-8.48l48 48a6 6 0 0 1 0 8.48l-48 48a6 6 0 0 1-8.48-8.48Z"/>`),
		g.Group(children),
	)
}