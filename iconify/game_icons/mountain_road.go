package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainRoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M149 24.32L123 124.4l27.8 44.8l15.9-29.4l37.5 30.9l17.5-55.5zm169.7 43.7l-32.2 48.38l10.7 25.5l21.9-21.7l18.1 42.8l10.3-21.7zm-83.4 64.18l-21.9 69.3l-41.3-33.9L152 205l-35.2-56.7l-88.38 339.4H154.9c45.6-17 135.3-56.7 137.2-106.3c1.7-47.8-108.9-43.4-110.8-91.2c-1.5-39.1 84.5-81.5 84.5-81.5s-54.1 43.3-47.4 71c12.9 53 125.6 27.6 143.2 79.3c13 38.2-33.5 104.3-52.2 128.7h174.2L356.5 164.3L336.4 207l-23.5-55.4l-22.1 21.9l-16.3-39.1l-17 25.5z"/>`),
		g.Group(children),
	)
}