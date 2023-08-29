package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 16.379l-6.173 3.245l1.179-6.874L.01 7.882l6.902-1.003L10 .624l3.087 6.255l6.902 1.003l-4.995 4.868l1.18 6.874z"/>`),
		g.Group(children),
	)
}