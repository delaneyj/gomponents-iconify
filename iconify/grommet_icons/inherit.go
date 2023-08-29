package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inherit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m17 18l-5-3l5 3ZM7 18l5-3v-4m5 9a3 3 0 1 0 6 0a3 3 0 0 0-6 0ZM4 17a3 3 0 1 0 0 6a3 3 0 0 0 0-6ZM17 6a5 5 0 1 1-10.001-.001A5 5 0 0 1 17 6Z"/>`),
		g.Group(children),
	)
}