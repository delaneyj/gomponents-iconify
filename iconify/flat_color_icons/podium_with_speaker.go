package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumWithSpeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="11" r="6" fill="#FFB74D"/><path fill="#607D8B" d="M36 26.1S32.7 19 24 19s-12 7.1-12 7.1V30h24v-3.9z"/><path fill="#B0BEC5" d="M41 25H7l-1 4l5 3l-2-3h30l-2 3l5-3z"/><path fill="#78909C" d="M9 29h30l-4 12H13z"/>`),
		g.Group(children),
	)
}