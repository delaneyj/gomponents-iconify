package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchAdvanced(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m15 16l6 6l-6-6Zm-5 2a7 7 0 1 0 0-14a7 7 0 0 0 0 14ZM20 1v6m-3-3h6"/>`),
		g.Group(children),
	)
}