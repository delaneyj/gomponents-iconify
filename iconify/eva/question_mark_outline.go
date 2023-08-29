package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionMarkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaQuestionMarkOutline0"><g id="evaQuestionMarkOutline1"><g id="evaQuestionMarkOutline2" fill="currentColor"><path d="M17 9A5 5 0 0 0 7 9a1 1 0 0 0 2 0a3 3 0 1 1 3 3a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-1.1A5 5 0 0 0 17 9Z"/><circle cx="12" cy="19" r="1"/></g></g></g>`),
		g.Group(children),
	)
}