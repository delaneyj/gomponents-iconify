package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 64a64 64 0 0 0-64 64a64 64 0 0 0 1.1 11.3l-146.3 73.2A64 64 0 0 0 128 192a64 64 0 0 0-64 64a64 64 0 0 0 64 64a64 64 0 0 0 46.8-20.5L321 372.7a64 64 0 0 0-1 11.3a64 64 0 0 0 64 64a64 64 0 0 0 64-64a64 64 0 0 0-64-64a64 64 0 0 0-46.8 20.5L191 267.4a64 64 0 0 0 1-11.4a64 64 0 0 0-1.1-11.4l146.3-73.1A64 64 0 0 0 384 192a64 64 0 0 0 64-64a64 64 0 0 0-64-64z"/>`),
		g.Group(children),
	)
}