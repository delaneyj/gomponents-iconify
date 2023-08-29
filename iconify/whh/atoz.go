package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atoz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M741 959h251q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5H672q-13 0-22.5-9t-9.5-23q0-6 4-12.5t6.5-8T661 961l263-322H672q-13 0-22.5-9t-9.5-22.5t9.5-23T672 575h320q11 0 21.5 9.5T1024 607q0 10-20 30zm251-511q-13 0-22.5-9.5T960 416v-80q0-7-4.5-11.5T944 320H720q-6 0-11 4.5t-5 11.5v80q0 13-9 22.5t-22.5 9.5t-23-9.5T640 416V128q0-53 37.5-90.5T768 0h128q53 0 90.5 37.5T1024 128v288q0 13-9.5 22.5T992 448zm-32-320q0-27-19-45.5T896 64H768q-26 0-45 18.5T704 128v112q0 6 5 11t11 5h224q7 0 11.5-5t4.5-11V128zm-742 887q-11 8-26 8t-26-8L11 808q-11-8-11-20t11-20h118V64q0-27 18.5-45.5t45-18.5T238 18.5T257 64v704h116q11 8 11 20t-11 20z"/>`),
		g.Group(children),
	)
}