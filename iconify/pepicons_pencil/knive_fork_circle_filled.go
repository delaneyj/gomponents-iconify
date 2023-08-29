package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KniveForkCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilKniveForkCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M17.28 4.051C15.12 5.115 14 6.611 14 8.506c0 1.82 1.037 3.051 3 3.631V14a.5.5 0 0 0 1 0V4.5a.5.5 0 0 0-.72-.449ZM15 8.506c0-1.268.65-2.314 2-3.158v5.74c-1.364-.477-2-1.321-2-2.582Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16 15.5v5a1.5 1.5 0 0 0 3 0v-5a1.5 1.5 0 0 0-3 0Zm1.5 5.5a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 1 0v5a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path d="M7.532 5.475a.5.5 0 1 1 1 .05c-.084 1.672.003 2.81.237 3.374c.183.442.544.634 1.202.597a.5.5 0 0 1 .058.998c-1.058.06-1.826-.349-2.184-1.211c-.307-.74-.403-1.996-.313-3.808Z"/><path d="M12.467 5.475a.5.5 0 1 0-.998.05c.083 1.672-.004 2.81-.238 3.374c-.183.442-.544.634-1.202.597a.5.5 0 0 0-.058.998c1.058.06 1.826-.349 2.184-1.211c.307-.74.403-1.996.312-3.808Z"/><path d="M9.5 5.5a.5.5 0 0 1 1 0v9a.5.5 0 0 1-1 0v-9Z"/><path fill-rule="evenodd" d="M8.5 15.5v5a1.5 1.5 0 0 0 3 0v-5a1.5 1.5 0 0 0-3 0ZM10 21a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 1 0v5a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilKniveForkCircleFilled0)"/></g>`),
		g.Group(children),
	)
}