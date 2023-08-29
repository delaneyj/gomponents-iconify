package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionCurveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm-5.414 11.414L18 8.828v5.769a5.02 5.02 0 0 1-1.096 3.124l-2.247 2.808A3.01 3.01 0 0 0 14 22.403V27h-2v-4.597a5.02 5.02 0 0 1 1.096-3.124l2.247-2.808A3.01 3.01 0 0 0 16 14.597V8.828l-4.586 4.586L10 12l7-7l7 7Z"/><path fill="none" d="M22.586 13.414L18 8.828v5.769a5.02 5.02 0 0 1-1.096 3.124l-2.247 2.808A3.01 3.01 0 0 0 14 22.403V27h-2v-4.597a5.02 5.02 0 0 1 1.096-3.124l2.247-2.808A3.01 3.01 0 0 0 16 14.597V8.828l-4.586 4.586L10 12l7-7l7 7Z"/>`),
		g.Group(children),
	)
}