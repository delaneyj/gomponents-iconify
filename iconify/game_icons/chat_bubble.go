package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M229.7 22.66A155.2 235.6 80.24 0 0 23.81 215.6A155.2 235.6 80.24 0 0 236.7 333.4c23.8 55.6-17.1 109.3-83.6 161.1c86.2-28.3 176.2-94.4 179.7-178.7a155.2 235.6 80.24 0 0 155.4-180.1A155.2 235.6 80.24 0 0 229.7 22.66z"/>`),
		g.Group(children),
	)
}