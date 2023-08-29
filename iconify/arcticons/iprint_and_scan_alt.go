package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IprintAndScanAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.997 15.118L8.86 8.425l29.875-.154l4.564 7.045Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.997 15.118V17.2l38.303.002v-1.887m-2.002 1.887L43.5 26.53l-17.12.1m-3.444 0H4.5l2.084-9.427M4.5 26.63l2.183 9.427l6.68.031m21.106.25l6.45-.082L43.5 26.53m-31.607-9.327l12.553 10.716l11.908-7.938L33 17.202M13.233 29.111l.198 10.617l21.038-.298V29.11"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.274 31.294H10.75l-.198-2.183h26.992l-.198 2.183h-2.878m-19.153 5.904L26.63 37m-11.314-2.44H26.63m-11.314-2.82l11.314-.099"/>`),
		g.Group(children),
	)
}