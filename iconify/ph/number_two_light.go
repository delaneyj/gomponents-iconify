package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M174 208a6 6 0 0 1-6 6H88a6 6 0 0 1-4.8-9.6l72-95.92a34 34 0 0 0-31.88-54.14a34.1 34.1 0 0 0-26.49 20.07a6 6 0 0 1-11-4.81a46.47 46.47 0 0 1 5.43-9.28a46 46 0 0 1 73.48 55.37L100 202h68a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}