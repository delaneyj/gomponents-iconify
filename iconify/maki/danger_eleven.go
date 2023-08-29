package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DangerEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6.62 8.5l3.11 1.55l-.45.89L5.5 9.06l-3.78 1.89l-.45-.89L4.38 8.5l-3.1-1.55l.45-.89L5.5 7.94l3.78-1.89l.44.9l-3.1 1.55zM8.5 3.21v.29l-1 1v1l-2 1l-2-1v-1l-1-1V3a2.9 2.9 0 0 1 3-3a3.09 3.09 0 0 1 3 3.21zm-3.79-.5a.79.79 0 1 0-.79.79a.79.79 0 0 0 .79-.79zM5 4.5h-.5v1H5v-1zm1.5 0H6v1h.5v-1zm1.36-1.79a.79.79 0 1 0-.79.79a.79.79 0 0 0 .79-.79z" fill="currentColor"/>`),
		g.Group(children),
	)
}