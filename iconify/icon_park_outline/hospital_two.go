package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 11h8l4 10H4l4-10h8M8 21h32v22H8V21Z"/><path d="M16 5h16v16H16zm0 24h8v14h-8zm8 0h8v14h-8zm-3-16h6m9 30H12m12-27v-6"/></g>`),
		g.Group(children),
	)
}