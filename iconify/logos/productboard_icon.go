package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductboardIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#FF2638" d="m85.327 83.997l85.327 83.996H0z"/><path fill="#FFC600" d="m0 0l85.327 83.997L170.654 0z"/><path fill="#0079F2" d="m85.341 83.997l85.327 83.996l85.327-83.996L170.668 0z"/>`),
		g.Group(children),
	)
}