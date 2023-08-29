package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPeas0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M27 12s7 7 7 16s-4.445 16.223-8 16c-3.556-.223-7-7-6-16s7-16 7-16Zm0 0s1-4.124 4-6.062C34 4 39.89 9 39 12c-.89 3-4 3-5 0s4-6.5 7-6.062c3 .438 3.257 5.242 3 8.062c-.501 5.5-2 10-2 10"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M27 12s0 8-7 16s-13.675 9.7-16 7c-2.325-2.7 0-10 7-17s16-6 16-6Z"/><circle cx="27.243" cy="27.408" r="2.5" fill="#fff"/><circle cx="26.243" cy="34.408" r="2.5" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPeas0)"/>`),
		g.Group(children),
	)
}