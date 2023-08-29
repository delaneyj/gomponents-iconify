package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnorderedList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 18v9m3 0H5m6-15H5m0-7s3-3 5 0s-5 7-5 7m0 22.5s2-3 5-1s0 4.5 0 4.5s3 2.5 0 4.5s-5-1-5-1m5-3.5H8m0-20l-3 2m15 4h22M20 38h22M20 10h22"/>`),
		g.Group(children),
	)
}