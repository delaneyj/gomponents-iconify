package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M23 23V11H9v12h14Zm-9-6h4m-7-6V7c0-3 0-6-5-6S1 4 1 7v4"/>`),
		g.Group(children),
	)
}