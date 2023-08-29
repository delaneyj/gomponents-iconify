package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 28V17h1a2.002 2.002 0 0 0 2-2v-4a2.002 2.002 0 0 0-2-2h-4a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h1v11H2v2h28v-2Zm-3-17h4l.002 4H14Z"/><path fill="currentColor" d="M9.332 18.217a7 7 0 0 1 0-10.434l1.334 1.49a5 5 0 0 0 0 7.453zm13.335 0l-1.334-1.49a5 5 0 0 0 0-7.454l1.334-1.49a7 7 0 0 1 0 10.434z"/><path fill="currentColor" d="M6.4 21.8a11.002 11.002 0 0 1 0-17.6l1.2 1.6a9 9 0 0 0 0 14.401zm19.2 0l-1.2-1.6a9.001 9.001 0 0 0 0-14.401l1.2-1.6a11.002 11.002 0 0 1 0 17.601z"/>`),
		g.Group(children),
	)
}