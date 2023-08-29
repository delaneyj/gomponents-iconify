package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Launch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 9v5s-3 2-3 5v1h4l2 3h4l2-3h4v-1c0-3-3-5-3-5V9c0-4-3-8-5-8S7 5 7 9Zm1 11h8M12 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}