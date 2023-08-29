package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulgaria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m49.66 101.8l40.48 16.1l-11.03 28.4l185.19 11.6l95.3-47.1L483 150.1l5 40.1l-38.9 3.3l-8.9 63.9S405.5 281 408.9 281c3.5 0 44.5 44.9 44.5 44.9l-40.8 16.5l-15.4-17.5l-89 39.3l5.6 30.4l-72.1 15.6l-59.9-34.4l-121 21.2l3.25-67.5l-40.03-25.6l15.44-63.3l41.15-31.8l-50.44-54.4z"/>`),
		g.Group(children),
	)
}