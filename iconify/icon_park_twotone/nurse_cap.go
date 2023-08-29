package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NurseCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTNurseCap0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 10c-11.257 0-18.015 6.748-19.625 8.558a.924.924 0 0 0-.148.982l7.326 17.397c.236.561.917.794 1.467.532C14.824 36.608 18.876 35 24 35c5.124 0 9.176 1.608 10.98 2.47c.55.261 1.231.028 1.467-.533l7.326-17.397a.924.924 0 0 0-.148-.982C42.015 16.748 35.257 10 24 10Z"/><path d="M20 24.001h8M24 20v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTNurseCap0)"/>`),
		g.Group(children),
	)
}