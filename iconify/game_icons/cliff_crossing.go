package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CliffCrossing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m194.8 103.5l-5.6 17c49.1 16.4 84.5 16.4 133.6 0l-5.6-17c-46.9 15.6-75.5 15.6-122.4 0zM25 121.9V487h158V339.7l-15.5 46.6l-23.3-99.7l34.4 17.8l-11.7-47.9l15-120.4l-96.71-8.8l-13.61 46.4L59.39 125L25 121.9zm462 0l-156.8 14.2l30.9 232.4L331.5 487h68.3l-12.2-48.7l39.2-3.7l-20.4 52.4H487v-86.2L466.1 328l20.9 1.4V121.9zm-16.8 102.2l-15.8 41.2l-26.5-11.2l-29.7 31.4l-19.2-44.9l91.2-16.5zM42.95 402.7l120.35 17.6l-32.9 41.3l-40.26-29.1l-21.74 47.7l-25.45-77.5z"/>`),
		g.Group(children),
	)
}