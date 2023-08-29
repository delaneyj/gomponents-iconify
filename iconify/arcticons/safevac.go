package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Safevac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.694 0 15.24-7.79 15.24-16.943V6.505C35.228 6.505 24 4.5 24 4.5S12.671 6.505 8.76 6.505v20.052C8.76 35.75 22.397 43.5 24 43.5Zm9.219-30.457l-2.497-2.176m.692 8.693l-7.008-6.126m7.208-1.094l-3.639 4.161"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.61 17.955l-9.405 10.728l-3.328-2.908l9.404-10.727M18.48 27.28l-3.699 4.22"/>`),
		g.Group(children),
	)
}