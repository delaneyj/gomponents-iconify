package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserBlockLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="6" r="4"/><path d="M15 13.327A13.57 13.57 0 0 0 12 13c-4.418 0-8 2.015-8 4.5S4 22 12 22c5.687 0 7.331-1.018 7.807-2.5" opacity=".5"/><path stroke-linejoin="round" d="M15.171 18.828a4 4 0 1 0 5.657-5.657m-5.656 5.657a4 4 0 0 1 5.657-5.657m-5.658 5.657l5.657-5.656"/></g>`),
		g.Group(children),
	)
}