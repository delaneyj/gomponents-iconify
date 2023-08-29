package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicePresentation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2196F3" d="M40 22h-8l-4 4V12c0-2.2 1.8-4 4-4h8c2.2 0 4 1.8 4 4v6c0 2.2-1.8 4-4 4z"/><circle cx="17" cy="19" r="8" fill="#FFA726"/><path fill="#607D8B" d="M30 36.7S26.4 30 17 30S4 36.7 4 36.7V40h26v-3.3z"/>`),
		g.Group(children),
	)
}