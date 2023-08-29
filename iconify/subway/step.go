package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Step(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M217 39.5C95.5 49.5 0 151.1 0 275.2C0 406 106 512 236.8 512c124.1 0 225.7-95.5 235.8-217H217V39.5zM256.5 0v255.5H512C502.3 118.7 393.3 9.7 256.5 0z"/>`),
		g.Group(children),
	)
}