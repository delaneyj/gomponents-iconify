package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Popcorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPopcorn0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M33.696 40.868L39 17H7l5.304 23.868c.334 1.501.5 2.252 1.049 2.692c.548.44 1.317.44 2.856.44H29.79c1.539 0 2.308 0 2.856-.44c.549-.44.715-1.19 1.05-2.692Z"/><path stroke="#000" d="m27 44l1-27m-9 27l-1-27"/><path stroke="#fff" d="M31 44H15m16-27H15m-4 0s-1-3 .5-4.5s4.5-1 4.5-1s0-3 2-4s5 .5 5 .5s2-3.357 5-2.5c3 .857 3 4.5 3 4.5s2.5 0 4 2s0 5 0 5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPopcorn0)"/>`),
		g.Group(children),
	)
}