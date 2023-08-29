package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="36" x="4" y="8" rx="2"/><path d="M4 20h40M4 32h40M17 4v8m14-8v8m-14 8v24m14-24v24m13-31v26M4 13v26m10 5h20"/></g>`),
		g.Group(children),
	)
}