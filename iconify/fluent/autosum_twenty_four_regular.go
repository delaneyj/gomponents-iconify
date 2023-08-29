package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutosumTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.06 4.457A.75.75 0 0 1 5.75 4h12.5a.75.75 0 0 1 0 1.5H7.52l5.36 5.539a.75.75 0 0 1 .03 1.01L7.38 18.5h10.87a.75.75 0 0 1 0 1.5H5.75a.75.75 0 0 1-.57-1.238l6.147-7.17l-6.116-6.32a.75.75 0 0 1-.152-.815Z"/>`),
		g.Group(children),
	)
}