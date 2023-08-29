package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NurseCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 10c-11.257 0-18.015 6.748-19.625 8.558a.924.924 0 0 0-.148.982l7.326 17.397c.236.561.917.794 1.467.532C14.824 36.608 18.876 35 24 35c5.124 0 9.176 1.608 10.98 2.47c.55.261 1.231.028 1.467-.533l7.326-17.397a.924.924 0 0 0-.148-.982C42.015 16.748 35.257 10 24 10Zm-4 14.001h8M24 20v8"/>`),
		g.Group(children),
	)
}