package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adnihilation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="24.285" x="4.5" y="11.857" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.057 15.15l23.161 17.01m-12.35-17.515l3.097 18.852m-14.703-7.328l25.001-3.563m-3.933-3.872L15.641 30.448"/>`),
		g.Group(children),
	)
}