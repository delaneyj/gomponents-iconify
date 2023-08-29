package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 4)"><path d="M2.727 7.033A7.539 7.539 0 0 1 5.492 5.61m4.05-.005a7.54 7.54 0 0 1 2.785 1.43M7.5 8.5l.027-8M.286 4.667A10.974 10.974 0 0 1 5.511 2.18m4.087.02a10.972 10.972 0 0 1 5.116 2.467m-9.58 4.74c.161-.112.328-.211.5-.298m3.706-.016c.183.09.361.195.533.314"/><circle cx="7.5" cy="11.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}