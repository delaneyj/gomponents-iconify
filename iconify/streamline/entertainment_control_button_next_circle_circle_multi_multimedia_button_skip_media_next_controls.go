package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentControlButtonNextCircleCircleMultiMultimediaButtonSkipMediaNextControls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M9.5 5v4m-5-4.5l3 2.5l-3 2.5v-5z"/></g>`),
		g.Group(children),
	)
}