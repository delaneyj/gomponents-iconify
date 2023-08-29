package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepsOption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M24 9h-5v5h-5v5H9v5m-8 0v-7.003c0-.55.313-1.31.703-1.7L15.297 1.703c.388-.388 1.156-.703 1.7-.703H24"/>`),
		g.Group(children),
	)
}