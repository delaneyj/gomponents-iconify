package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperBagFolded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m151.7 112.2l-23.3 42.7l16.2 8.9l172.8 13.5l28.7-23l-32.2 5.3l14.3-35.9l-176.5-11.5zM346 125.8l21.8 27.3l-9.9 102.9l2.8 120.1l43.6 43.3l-12.1-51.2l5.5-112.9l-25.3-109.7l-26.4-19.8zm-1.2 47.1l-34.6 29.4l-162.4-14.7l-33.3 44.2l-13.4 124.6l173.2 28.1l-174.88-12.4l-7.28 67.7l195.26 26.7l8.8-216.5l48.6-77.1zm6.1 18.8L313.7 256l-6.2 185.9l35.9-60.6l.3-125.3l7.2-64.3zm3.6 201.1l-47 79.7l92.8-30.6l-45.8-49.1z"/>`),
		g.Group(children),
	)
}