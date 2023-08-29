package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="M7 41V23l18 18z"/><path fill="#3F51B5" d="m35.355 6.946l5.656 5.656l-23.119 23.12l-5.656-5.657z"/>`),
		g.Group(children),
	)
}