package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="M41 7v18L23 7z"/><path fill="#3F51B5" d="m12.641 40.983l-5.656-5.656l23.12-23.119l5.655 5.656z"/>`),
		g.Group(children),
	)
}