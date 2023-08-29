package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FollowUpDateSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 5V30.0036H42V5"/><path stroke-linejoin="round" d="M30 37L24 43L18 37"/><path stroke-linejoin="round" d="M24 30V43"/><path stroke-linejoin="round" d="M27.9887 10.9785L33 16L27.9887 21.0902"/><path d="M15.0009 16.001H33"/></g>`),
		g.Group(children),
	)
}