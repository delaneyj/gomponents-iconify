package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.5 17H1V1h16v3.5M7 7h16v16H7V7Zm8 4v8v-8Zm-4 4h8h-8Z"/>`),
		g.Group(children),
	)
}