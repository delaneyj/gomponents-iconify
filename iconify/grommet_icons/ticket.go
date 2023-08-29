package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 16h10V8H7v8Zm13-4c0 2 1 3 3 3v5H1v-5c2 0 3-1 3-3S3 9 1 9V4h22v5c-2 0-3 1-3 3Z"/>`),
		g.Group(children),
	)
}