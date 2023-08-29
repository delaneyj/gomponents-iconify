package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="M7 7h18L7 25z"/><path fill="#3F51B5" d="m41.02 35.322l-5.656 5.656l-23.12-23.119l5.657-5.656z"/>`),
		g.Group(children),
	)
}