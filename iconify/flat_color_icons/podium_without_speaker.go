package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumWithoutSpeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#B0BEC5" d="M43 16H5l-1 4l5 3l-2-3h34l-2 3l5-3z"/><path fill="#78909C" d="M7 20h34l-4 16H11z"/>`),
		g.Group(children),
	)
}