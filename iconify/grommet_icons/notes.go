package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 1v22h13l5-5V1H3Zm3 16h5m-5-4h12M6 9h10M3 5h18m0 12h-6v6"/>`),
		g.Group(children),
	)
}