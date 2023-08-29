package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChristmasTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSChristmasTree0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="m20 14l-4-2l8-8l8 8l-4 2l8 8l-6 2l9 10H9l8-10l-5-2l8-8Z"/><path d="M31 44H17m4-10l-1 10m7-10l1 10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSChristmasTree0)"/>`),
		g.Group(children),
	)
}