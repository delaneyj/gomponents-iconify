package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvShowTrackerTrakttvClient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.809 13.75L24 3.5L6.191 13.75v20.5L24 44.5l17.809-10.25v-20.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.456 21.694a2.57 2.57 0 0 1-2.562-2.563c0-.897.512-1.665 1.153-2.178c-1.794-1.794-4.356-2.947-7.047-2.947c-5.51 0-9.994 4.485-9.994 9.994s4.485 9.994 9.994 9.994s9.994-4.485 9.994-9.994c0-.897-.128-1.794-.385-2.69c-.256.256-.769.384-1.153.384Z"/>`),
		g.Group(children),
	)
}