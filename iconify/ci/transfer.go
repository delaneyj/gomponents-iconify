package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 20l-5-4l5-4v3h13v2H9v3Zm6-8V9H2V7h13V4l5 4l-5 4Z"/>`),
		g.Group(children),
	)
}