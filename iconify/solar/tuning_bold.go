package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.959 9.75a.75.75 0 0 1-.75-.75V2a.75.75 0 0 1 1.5 0v7a.75.75 0 0 1-.75.75ZM7 12.5a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm10-1a3 3 0 1 0 0 6a3 3 0 0 0 0-6ZM6.209 15a.75.75 0 0 1 1.5 0v7a.75.75 0 0 1-1.5 0v-7Zm10.75 7.75a.75.75 0 0 1-.75-.75v-2a.75.75 0 0 1 1.5 0v2a.75.75 0 0 1-.75.75ZM6.209 2a.75.75 0 0 1 1.5 0v2a.75.75 0 0 1-1.5 0V2Z"/>`),
		g.Group(children),
	)
}