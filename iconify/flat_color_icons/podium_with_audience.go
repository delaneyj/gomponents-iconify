package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumWithAudience(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#B0BEC5" d="M41 12H7l-1 4l5 3l-2-3h30l-2 3l5-3z"/><path fill="#78909C" d="M9 16h30l-4 12H13z"/><circle cx="24" cy="28" r="4" fill="#FFB74D"/><circle cx="36" cy="28" r="4" fill="#FFB74D"/><circle cx="12" cy="28" r="4" fill="#FFB74D"/><circle cx="18" cy="37" r="5" fill="#FFB74D"/><circle cx="30" cy="37" r="5" fill="#FFB74D"/>`),
		g.Group(children),
	)
}