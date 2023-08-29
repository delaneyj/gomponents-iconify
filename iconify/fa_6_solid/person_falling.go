package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonFalling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M288 0c17.7 0 32 14.3 32 32v9.8c0 54.6-27.9 104.6-72.5 133.6l.2.3l56.8 80.3H392c15.1 0 29.3 7.1 38.4 19.2l43.2 57.6c10.6 14.1 7.7 34.2-6.4 44.8s-34.2 7.7-44.8-6.4L384 320h-97.4l92.3 142.6c9.6 14.8 5.4 34.6-9.5 44.3s-34.6 5.4-44.3-9.5L164.5 249.2c-2.9 9.2-4.5 19-4.5 29V352c0 17.7-14.3 32-32 32s-32-14.3-32-32v-73.8c0-65.1 39.6-123.7 100.1-147.9c36.2-14.5 59.9-49.5 59.9-88.5V32c0-17.7 14.3-32 32-32zM112 32a48 48 0 1 1 0 96a48 48 0 1 1 0-96z"/>`),
		g.Group(children),
	)
}