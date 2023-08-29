package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficSignLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M241.74 117.71L138.29 14.26a14.56 14.56 0 0 0-20.58 0L14.26 117.71a14.56 14.56 0 0 0 0 20.58l103.45 103.45a14.56 14.56 0 0 0 20.58 0l103.45-103.45a14.56 14.56 0 0 0 0-20.58Zm-8.49 12.1L129.81 233.25a2.56 2.56 0 0 1-3.62 0L22.75 129.81a2.56 2.56 0 0 1 0-3.62L126.19 22.75a2.56 2.56 0 0 1 3.62 0l103.44 103.44a2.56 2.56 0 0 1 0 3.62Zm-61-14a6 6 0 0 1 0 8.48l-24 24a6 6 0 0 1-8.48-8.48L153.51 126H112a18 18 0 0 0-18 18v8a6 6 0 0 1-12 0v-8a30 30 0 0 1 30-30h41.51l-13.75-13.76a6 6 0 0 1 8.48-8.48Z"/>`),
		g.Group(children),
	)
}