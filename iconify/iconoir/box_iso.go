package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxIso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-width="1.5"><path fill="currentColor" d="m2.695 7.185l9 4l.61-1.37l-9-4l-.61 1.37ZM12.75 21.5v-11h-1.5v11h1.5Zm-.445-10.315l9-4l-.61-1.37l-9 4l.61 1.37Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 17.11V6.89a.6.6 0 0 1 .356-.548l8.4-3.734a.6.6 0 0 1 .488 0l8.4 3.734A.6.6 0 0 1 21 6.89v10.22a.6.6 0 0 1-.356.548l-8.4 3.734a.6.6 0 0 1-.488 0l-8.4-3.734A.6.6 0 0 1 3 17.11Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 4.5l8.644 3.842a.6.6 0 0 1 .356.548v3.61"/></g>`),
		g.Group(children),
	)
}