package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizDistributionRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 17V7m0 10h5.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H5m0 10v3M5 7V4m10 13V7m0 10h3.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H15m0 10v3m0-13V4"/>`),
		g.Group(children),
	)
}