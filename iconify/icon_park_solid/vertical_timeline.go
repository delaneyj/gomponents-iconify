package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalTimeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSVerticalTimeline0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M4 5h40M4 43h40M8 36v7"/><path fill="#fff" stroke-linejoin="round" d="M12 28H4v8h8v-8Zm16-8h-8v8h8v-8Zm16-8h-8v8h8v-8Z"/><path stroke-linecap="round" d="M40 20v23M8 12v1m0 7v1m15-9v1m1 15v15"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSVerticalTimeline0)"/>`),
		g.Group(children),
	)
}