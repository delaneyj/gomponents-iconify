package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cmake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.693.088L.088 30.943l17.016-14.459zm15.25 31.824L9.964 23.448L0 31.912zM32 31.645L16.396.62l2.292 25.651zM17.193 17.281l-6.704 5.729l7.495 2.995z"/>`),
		g.Group(children),
	)
}