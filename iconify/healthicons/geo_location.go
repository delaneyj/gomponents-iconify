package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GeoLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M23.424 41.817L24 41l-.576.817Zm1.152 0l.004-.002l.01-.007l.03-.023l.118-.085c.1-.074.246-.182.43-.324c.368-.282.89-.697 1.513-1.23a43.403 43.403 0 0 0 4.575-4.54C34.564 31.78 38 26.319 38 20.076c0-3.73-1.474-7.31-4.098-9.95A13.962 13.962 0 0 0 24 6a13.962 13.962 0 0 0-9.902 4.125A14.117 14.117 0 0 0 10 20.077c0 6.242 3.436 11.703 6.744 15.529a43.403 43.403 0 0 0 4.575 4.54c.624.533 1.145.948 1.513 1.23a25.536 25.536 0 0 0 .547.41l.032.022l.009.007l.004.002c.345.243.807.243 1.152 0ZM24 41l.576.817L24 41Zm5-21a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}