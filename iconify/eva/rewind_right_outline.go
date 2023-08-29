package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaRewindRightOutline0"><g id="evaRewindRightOutline1"><path id="evaRewindRightOutline2" fill="currentColor" d="m20.86 10.67l-5.1-4.21a2.1 2.1 0 0 0-2.21-.26a1.76 1.76 0 0 0-1.05 1.59v2.59L7.76 6.46a2.1 2.1 0 0 0-2.21-.26a1.76 1.76 0 0 0-1 1.59v8.42a1.76 1.76 0 0 0 1 1.59a2.23 2.23 0 0 0 .91.2a2.06 2.06 0 0 0 1.3-.46l4.74-3.92v2.59a1.76 1.76 0 0 0 1.05 1.59a2.23 2.23 0 0 0 .91.2a2.06 2.06 0 0 0 1.3-.46l5.1-4.21a1.7 1.7 0 0 0 0-2.66ZM6.5 15.91V8l4.82 4Zm8 0V8l4.82 4Z"/></g></g>`),
		g.Group(children),
	)
}