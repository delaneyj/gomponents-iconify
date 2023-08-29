package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbbotMeeple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 16.1l-82.7 82.72l35 69.98a66.18 66.18 0 0 0-18.6 45.8a66.18 66.18 0 0 0 10.4 35.3c-42.8 5.3-126.12 21.2-126.12 47.5c0 16.8 30.02 37.8 68.12 54.5L73.98 462.8c0 33.1 33.02 33.1 66.22 33.1h231.5c33.1 0 66.3 0 66.3-33.1l-68.2-110.9c38.1-16.7 68.2-37.7 68.2-54.5c0-26.3-83.5-42.2-126.1-47.5a66.18 66.18 0 0 0 10.2-35.3a66.18 66.18 0 0 0-18.4-45.8l35-69.98z"/>`),
		g.Group(children),
	)
}