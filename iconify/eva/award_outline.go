package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaAwardOutline0"><g id="evaAwardOutline1"><path id="evaAwardOutline2" fill="currentColor" d="m19 20.75l-2.31-9A5.94 5.94 0 0 0 18 8A6 6 0 0 0 6 8a5.94 5.94 0 0 0 1.34 3.77L5 20.75a1 1 0 0 0 1.48 1.11l5.33-3.13l5.68 3.14A.91.91 0 0 0 18 22a1 1 0 0 0 1-1.25ZM12 4a4 4 0 1 1-4 4a4 4 0 0 1 4-4Zm.31 12.71a1 1 0 0 0-1 0l-3.75 2.2L9 13.21a5.94 5.94 0 0 0 5.92 0L16.45 19Z"/></g></g>`),
		g.Group(children),
	)
}