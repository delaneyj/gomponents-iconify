package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M391 256c0-35.5-21.7-65.9-52.5-78.7l-16.4 39.4c15.4 6.4 26.2 21.6 26.2 39.4c0 17.7-10.8 32.9-26.2 39.4l16.4 39.4c30.8-13 52.5-43.4 52.5-78.9zM371.3 98.5l-16.4 39.4c46.3 19.3 78.8 64.9 78.8 118.1c0 53.3-32.5 98.8-78.8 118.1l16.4 39.4c61.7-25.7 105-86.5 105-157.5S433 124.2 371.3 98.5zM49.7 149.3v213.3H135L284.3 512V0L135 149.3H49.7z"/>`),
		g.Group(children),
	)
}