package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M5 12a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm4 6v-2c0-2.25-1.787-4-4.036-4h.054C2.768 12 1 13.75 1 16v2M12 7h12M12 17h10m-10-5h7"/>`),
		g.Group(children),
	)
}