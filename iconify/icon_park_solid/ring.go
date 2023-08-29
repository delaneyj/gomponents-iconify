package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRing0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M24 11.619c2.093 0 3.79-1.706 3.79-3.81C27.79 5.707 26.093 4 24 4s-3.79 1.706-3.79 3.81c0 2.103 1.697 3.809 3.79 3.809ZM9.79 40.19c2.092 0 3.789-1.705 3.789-3.809s-1.697-3.81-3.79-3.81C7.697 32.572 6 34.278 6 36.382s1.697 3.81 3.79 3.81Zm28.42 0c2.093 0 3.79-1.705 3.79-3.809s-1.697-3.81-3.79-3.81s-3.79 1.706-3.79 3.81s1.697 3.81 3.79 3.81Z"/><path stroke-linecap="round" d="M33.143 10.314A18.105 18.105 0 0 1 42 25.904c0 .578-.027 1.148-.08 1.711v0m-10.906 14.96A17.863 17.863 0 0 1 24 44c-2.488 0-4.858-.507-7.014-1.425M6.08 27.615a18.416 18.416 0 0 1-.08-1.71a18.105 18.105 0 0 1 8.857-15.59"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRing0)"/>`),
		g.Group(children),
	)
}