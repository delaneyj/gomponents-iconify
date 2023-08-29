package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="M41 41H23l18-18z"/><path fill="#3F51B5" d="m6.983 12.607l5.656-5.656l23.119 23.12l-5.656 5.655z"/>`),
		g.Group(children),
	)
}