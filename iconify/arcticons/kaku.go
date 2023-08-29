package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kaku(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.46 11.13V42.5m33.08-31.37V42.5m0-3.65H7.46M40.54 5.5H7.46m7.16 7.63h18.75v18.75H14.62zM24 5.5v26.38m-9.38-9.37h18.76"/>`),
		g.Group(children),
	)
}