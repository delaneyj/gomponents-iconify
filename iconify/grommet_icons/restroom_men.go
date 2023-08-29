package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestroomMen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 13.5L11 8l-1 13m7-7.5L13 8l1 13M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-1 3h2v5.5h-2V8Z"/>`),
		g.Group(children),
	)
}