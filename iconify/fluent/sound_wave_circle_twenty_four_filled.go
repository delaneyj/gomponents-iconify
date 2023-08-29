package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundWaveCircleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2Zm-1 6.75v6.5a.75.75 0 0 1-1.5 0v-6.5a.75.75 0 0 1 1.5 0Zm6 0v6.5a.75.75 0 0 1-1.5 0v-6.5a.75.75 0 0 1 1.5 0Zm-3 1.5v3.5a.75.75 0 0 1-1.5 0v-3.5a.75.75 0 0 1 1.5 0Zm-6 .5v2.5a.75.75 0 0 1-1.5 0v-2.5a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}