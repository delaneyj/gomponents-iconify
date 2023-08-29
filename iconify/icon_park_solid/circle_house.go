package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCircleHouse0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 23H4s10.5-6 15-11s5.5-8 5.5-8s1 3 5.5 8s14 11 14 11Z"/><path fill="#fff" d="M8 31h32v13H8zm5-8h22v8H13z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCircleHouse0)"/>`),
		g.Group(children),
	)
}