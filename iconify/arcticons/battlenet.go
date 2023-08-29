package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Battlenet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.312 27.494c-2.38 3.967-2.89 9.86-.453 11.729s8.386 2.096 19.264-10.652a28.408 28.408 0 0 0 7.083-20.284"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.08 32.942c4.625.078 9.983-2.427 10.384-5.472s-2.377-8.31-18.857-11.357A28.408 28.408 0 0 0 3.5 20.122"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.414 12.232c-2.245-4.043-7.093-7.431-9.93-6.256s-6.008 6.214-.407 22.009A28.408 28.408 0 0 0 31.102 44.26"/>`),
		g.Group(children),
	)
}