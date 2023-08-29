package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Campground(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25 6c-1.263 0-1.834.616-2.21 1.286L7.08 36.858H3.685C2.203 36.858 1 38.003 1 39.429c0 1.423 1.203 2.57 2.685 2.57h42.632c1.481 0 2.685-1.147 2.685-2.57c0-1.427-1.203-2.571-2.685-2.571H42.92L27.211 7.285C26.834 6.616 26.264 6 25 6zm0 12.857l8.842 18.001H16.158L25 18.857z"/>`),
		g.Group(children),
	)
}