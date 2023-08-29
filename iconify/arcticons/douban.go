package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Douban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.45 42.86h37m-37-35.8h37"/><rect width="31.53" height="17.1" x="8.18" y="12.93" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.37 30.03l-3.98 12.83m-13.4-7.97l2.47 7.97"/>`),
		g.Group(children),
	)
}