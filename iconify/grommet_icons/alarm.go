package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M21 13a9 9 0 0 1-9 9a9 9 0 0 1-9-9a9 9 0 0 1 9-9a9 9 0 0 1 9 9ZM5.5 19.5L2 23l3.5-3.5Zm13 0L22 23l-3.5-3.5ZM9 4c-.71-1.092-2.118-2-4-2c-2.1 0-4 1.9-4 4c0 1.882.908 3.29 2 4m18 0c1.092-.71 2-2.118 2-4c0-2.1-1.9-4-4-4c-1.882 0-3.29.908-4 2m-3 4v5l3 3"/>`),
		g.Group(children),
	)
}