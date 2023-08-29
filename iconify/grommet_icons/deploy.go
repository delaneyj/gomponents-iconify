package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deploy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M23 1s-6.528-.458-9 2c-.023.037-4 4-4 4L5 8l-3 2l8 4l4 8l2-3l1-5s3.963-3.977 4-4c2.458-2.472 2-9 2-9Zm-6 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM7 17c-1-1-3-1-4 0s-1 5-1 5s4 0 5-1s1-3 0-4Z"/>`),
		g.Group(children),
	)
}