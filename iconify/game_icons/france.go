package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func France(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M283.4 19.83c-3.2 0-31.2 5.09-31.2 5.09c-1.3 41.61-30.4 78.48-90.3 84.88l-12.8-23.07l-25.1 2.48l11.3 60.09l-113.79-4.9l12.2 41.5C156.3 225.4 150.7 338.4 124 439.4c47 53 141.8 47.8 186 43.1c3.1-62.2 52.4-64.5 135.9-32.2c11.3-17.6 18.8-36 44.6-50.7l-46.6-139.5l-27.5 6.2c11-21.1 32.2-49.9 50.4-63.4l15.6-86.9c-88.6-6.3-146.4-46.36-199-96.17z"/>`),
		g.Group(children),
	)
}