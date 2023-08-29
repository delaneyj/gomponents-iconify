package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireStationEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 0l-2 4L2 2c-.405.712-2 2.167-2 4c0 2.7 2.8 5 5.5 5S11 8.7 11 6c0-1.833-1.595-3.288-2-4L7.5 4l-2-4zm0 5.5s2 1.585 2 3c0 .611-.778 1.278-2 1.278s-2-.667-2-1.278c0-1.366 2-3 2-3z" fill="currentColor"/>`),
		g.Group(children),
	)
}