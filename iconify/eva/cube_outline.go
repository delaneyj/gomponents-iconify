package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCubeOutline0"><g id="evaCubeOutline1"><path id="evaCubeOutline2" fill="currentColor" d="M20.66 7.26c0-.07-.1-.14-.15-.21l-.09-.1a2.5 2.5 0 0 0-.86-.68l-6.4-3a2.7 2.7 0 0 0-2.26 0l-6.4 3a2.6 2.6 0 0 0-.86.68L3.52 7a1 1 0 0 0-.15.2A2.39 2.39 0 0 0 3 8.46v7.06a2.49 2.49 0 0 0 1.46 2.26l6.4 3a2.7 2.7 0 0 0 2.27 0l6.4-3A2.49 2.49 0 0 0 21 15.54V8.46a2.39 2.39 0 0 0-.34-1.2Zm-8.95-2.2a.73.73 0 0 1 .58 0l5.33 2.48L12 10.15L6.38 7.54ZM5.3 16a.47.47 0 0 1-.3-.43V9.1l6 2.79v6.72Zm13.39 0L13 18.61v-6.72l6-2.79v6.44a.48.48 0 0 1-.31.46Z"/></g></g>`),
		g.Group(children),
	)
}