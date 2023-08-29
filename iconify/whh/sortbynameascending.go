package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sortbynameascending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1003 61L741 384h251q13 0 22.5 9t9.5 22.5t-9.5 22.5t-22.5 9H672q-13 0-22.5-8.5T640 415q0-6 4-12t6.5-8t10.5-9L924 64H672q-13 0-22.5-9.5T640 32t9.5-22.5T672 0h320q11 0 21.5 9t10.5 22q0 10-21 30zm-785 953q-11 9-26 9t-26-9L11 808q-11-8-11-20t11-20h118V64q0-27 18.5-45.5t45-18.5T238 18.5T257 64v704h117q10 8 10 20t-10 20zm550-439h128q53 0 90.5 37.5T1024 703v288q0 13-9.5 22.5T992 1023t-22.5-9.5T960 991v-80q0-7-4.5-11.5T944 895H720q-6 0-11 4.5t-5 11.5v80q0 13-9 22.5t-22.5 9.5t-23-9.5T640 991V703q0-53 37.5-90.5T768 575zm-64 240q0 7 5 11.5t11 4.5h224q7 0 11.5-4.5T960 815V703q0-26-19-45t-45-19H768q-26 0-45 19t-19 45v112z"/>`),
		g.Group(children),
	)
}