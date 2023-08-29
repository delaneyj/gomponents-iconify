package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 11.5h-3v-5a.5.5 0 0 0-1 0V12a.5.5 0 0 0 .5.5h3.5a.5.5 0 0 0 0-1zM12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10zm0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9z"/>`),
		g.Group(children),
	)
}