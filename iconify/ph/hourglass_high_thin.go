package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassHighThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 28H72a12 12 0 0 0-12 12v36a12 12 0 0 0 4.8 9.6l56.53 42.4l-56.53 42.4A12 12 0 0 0 60 180v36a12 12 0 0 0 12 12h112a12 12 0 0 0 12-12v-35.64a12.05 12.05 0 0 0-4.76-9.57L134.63 128l56.61-42.79a12.05 12.05 0 0 0 4.76-9.57V40a12 12 0 0 0-12-12ZM72 36h112a4 4 0 0 1 4 4v20H68V40a4 4 0 0 1 4-4Zm116 144.36V216a4 4 0 0 1-4 4H72a4 4 0 0 1-4-4v-36a4 4 0 0 1 1.6-3.2L128 133l58.42 44.16a4 4 0 0 1 1.58 3.2Zm-1.59-101.53L128 123L69.6 79.2A4 4 0 0 1 68 76v-8h120v7.64a4 4 0 0 1-1.59 3.19Z"/>`),
		g.Group(children),
	)
}