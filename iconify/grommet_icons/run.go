package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Run(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l3 2m0-8a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM9.5 9.5L9.525 6H15L8 17H4m11-9l-3 5l.5 1L17 7.5L15 6m-4 7l5 3.5v5"/>`),
		g.Group(children),
	)
}