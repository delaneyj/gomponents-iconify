package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaRewindLeftOutline0"><g id="evaRewindLeftOutline1"><path id="evaRewindLeftOutline2" fill="currentColor" d="M18.45 6.2a2.1 2.1 0 0 0-2.21.26l-4.74 3.92V7.79a1.76 1.76 0 0 0-1.05-1.59a2.1 2.1 0 0 0-2.21.26l-5.1 4.21a1.7 1.7 0 0 0 0 2.66l5.1 4.21a2.06 2.06 0 0 0 1.3.46a2.23 2.23 0 0 0 .91-.2a1.76 1.76 0 0 0 1.05-1.59v-2.59l4.74 3.92a2.06 2.06 0 0 0 1.3.46a2.23 2.23 0 0 0 .91-.2a1.76 1.76 0 0 0 1.05-1.59V7.79a1.76 1.76 0 0 0-1.05-1.59ZM9.5 16l-4.82-4L9.5 8.09Zm8 0l-4.82-4l4.82-3.91Z"/></g></g>`),
		g.Group(children),
	)
}