package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaAttachOutline0"><g id="evaAttachOutline1"><path id="evaAttachOutline2" fill="currentColor" d="M9.29 21a6.23 6.23 0 0 1-4.43-1.88a6 6 0 0 1-.22-8.49L12 3.2A4.11 4.11 0 0 1 15 2a4.48 4.48 0 0 1 3.19 1.35a4.36 4.36 0 0 1 .15 6.13l-7.4 7.43a2.54 2.54 0 0 1-1.81.75a2.72 2.72 0 0 1-1.95-.82a2.68 2.68 0 0 1-.08-3.77l6.83-6.86a1 1 0 0 1 1.37 1.41l-6.83 6.86a.68.68 0 0 0 .08.95a.78.78 0 0 0 .53.23a.56.56 0 0 0 .4-.16l7.39-7.43a2.36 2.36 0 0 0-.15-3.31a2.38 2.38 0 0 0-3.27-.15L6.06 12a4 4 0 0 0 .22 5.67a4.22 4.22 0 0 0 3 1.29a3.67 3.67 0 0 0 2.61-1.06l7.39-7.43a1 1 0 1 1 1.42 1.41l-7.39 7.43A5.65 5.65 0 0 1 9.29 21Z"/></g></g>`),
		g.Group(children),
	)
}