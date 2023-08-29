package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoosterCarSeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBoosterCarSeat0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M32 15H16c0 3.47-5.78 7.903-9.317 10.291C5.123 26.345 4 27 4 27s1 3 3.5 8c1.894 3.788 4.075 5.854 5.02 6.633c.306.25.69.367 1.085.367h21.693c.455 0 .894-.154 1.217-.474c.86-.852 2.672-2.9 4.485-6.526c2.5-5 3-8 3-8s-1.278-.639-3-1.709c-3.554-2.207-9-6.25-9-10.291Z"/><path stroke="#fff" d="m32 15l2-9h8m-26 9l-2-9H6"/><path stroke="#000" d="M4 27s2 4 8 4h24c5 0 8-4 8-4"/><path stroke="#fff" d="M7.5 35C5 30 4 27 4 27s1.124-.655 2.683-1.709M41 35c2.5-5 3-8 3-8s-1.278-.639-3-1.709"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBoosterCarSeat0)"/>`),
		g.Group(children),
	)
}