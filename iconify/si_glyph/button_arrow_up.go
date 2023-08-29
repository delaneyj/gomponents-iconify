package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.984.057c4.389 0 7.959 3.562 7.959 7.941c0 4.377-3.57 7.939-7.959 7.939c-4.389 0-7.959-3.562-7.959-7.939c0-4.379 3.571-7.941 7.959-7.941zm-.015 14.026c3.347 0 6.074-2.729 6.074-6.083c0-3.353-2.727-6.083-6.074-6.083c-3.349 0-6.073 2.73-6.073 6.083c0 3.354 2.724 6.083 6.073 6.083z"/><path d="M11.975 7.906L8.973 4.002L6.022 7.907h2.003v3.264c0 .344.324.79.955.79c.63 0 .975-.483.975-.826V7.906h2.02z"/></g>`),
		g.Group(children),
	)
}