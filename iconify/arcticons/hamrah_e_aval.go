package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HamrahEAval(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.127 23.768c6.52-5.084 10.442-10.125 9.116-12.107c-1.589-2.372-10.103.543-19.017 6.51c-8.914 5.969-14.853 12.73-13.264 15.103c1.143 1.708 5.875.676 11.761-2.262"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.27 28.8c-5.06 2.234-8.178 4.665-7.727 6.389c.697 2.664 9.652 2.63 20.001-.077c10.35-2.707 18.175-7.06 17.479-9.725c-.382-1.461-3.251-2.11-7.473-1.952"/>`),
		g.Group(children),
	)
}