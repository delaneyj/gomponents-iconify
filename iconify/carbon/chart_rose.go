package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 17A15.017 15.017 0 0 0 15 2h-1v6.06A9.012 9.012 0 0 0 6.06 16H4v1a11.01 11.01 0 0 0 10 10.95V30h1a13.008 13.008 0 0 0 12.95-12H30ZM16 4.038A13.022 13.022 0 0 1 27.962 16H16Zm-2 6.042V16H8.08A7.004 7.004 0 0 1 14 10.08ZM6.055 18H14v7.939A9.013 9.013 0 0 1 6.055 18ZM16 27.955V18h9.955A11.02 11.02 0 0 1 16 27.955Z"/>`),
		g.Group(children),
	)
}