package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockClockwiseLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M134 80v44.6l37.09 22.25a6 6 0 0 1-6.18 10.3l-40-24A6 6 0 0 1 122 128V80a6 6 0 0 1 12 0Zm90-22a6 6 0 0 0-6 6v23.36c-7.48-8.83-14.94-17.13-23.53-25.83a94 94 0 1 0-1.95 134.83a6 6 0 0 0-8.24-8.72A82 82 0 1 1 186 70c9.24 9.36 17.18 18.3 25.31 28H184a6 6 0 0 0 0 12h40a6 6 0 0 0 6-6V64a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}