package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.68 2.12L12 8.8V2h6c.24 0 .46.05.68.12M9.5 7.5L7 9V2H6a2 2 0 0 0-2 2v12.8l7.88-7.87L9.5 7.5m12.11-5.77L1.89 21.46l1.27 1.27l1.38-1.38c.36.4.88.65 1.46.65h12c1.11 0 2-.89 2-2V5.89L22.89 3l-1.28-1.27Z"/>`),
		g.Group(children),
	)
}