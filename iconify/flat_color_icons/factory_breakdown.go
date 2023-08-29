package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactoryBreakdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#E64A19" d="M29 23v-6l-8 4v2h-4v4h-4v-4H5v20h28V23z"/><path fill="#992B0A" d="M25 27h4v4h-4zM9 35h4v4H9zm16 0h4v4h-4zm-8 0h4v4h-4zm0-8h4v4h-4z"/><path fill="#BF360C" d="M41.2 5H38v2h-2v2h-2.3L32 43h11z"/>`),
		g.Group(children),
	)
}