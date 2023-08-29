package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackSea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m416.5 116l-171.3 48.6l34.8 37.3l42.2-7l-.9 17.3l-28.4-2.4l-48.7 31.3l-19.3-10.4l8.9-22.4l-33.8-15.9l38.8-18.3l-4.5-11.4l-46.7-13l.5-29.8C109.9 155.5 72.82 236.6 25.74 305.2c-8.41 12.3 7.05 71.8 108.76 74c0 0 61.2-37.6 95.3-37.3c39.8.3 110.4 45.3 110.4 45.3l72.3 7.8c81.8 8.8 83.7-47.3 68.1-77.4c-7.5-14.5-91.7-76.9-147-104.3l-1.6-.1l4.1-6.4l22.5-4.3c4.2-13.1 3.8-28 28.1-32.9l-22.8-21.3l47.7-17.9z"/>`),
		g.Group(children),
	)
}