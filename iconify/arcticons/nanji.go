package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nanji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 42.462V7.91h9.049v31.02H5.5m0-15.095h9.049m5.57-12.146h19.906m-21.041 8.589h22.278m-23.515 7.287H42.5M29.473 5.538v14.74m6.323 0v18.065c0 2.278-.554 4.119-3.114 4.119H27.04m-5.65-10.141a13.927 13.927 0 0 1 5.472 5.255"/>`),
		g.Group(children),
	)
}