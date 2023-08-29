package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KniveFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M14.28 1.051C12.12 2.115 11 3.611 11 5.506c0 1.82 1.037 3.051 3 3.631V11a.5.5 0 0 0 1 0V1.5a.5.5 0 0 0-.72-.449ZM12 5.506c0-1.268.65-2.314 2-3.158v5.74c-1.364-.477-2-1.321-2-2.582Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 12.5v5a1.5 1.5 0 0 0 3 0v-5a1.5 1.5 0 0 0-3 0Zm1.5 5.5a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 1 0v5a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path d="M4.532 2.475a.5.5 0 1 1 1 .05c-.084 1.672.003 2.81.237 3.374c.183.442.544.634 1.202.597a.5.5 0 0 1 .058.998c-1.058.06-1.826-.349-2.184-1.211c-.307-.74-.403-1.996-.313-3.808Z"/><path d="M9.467 2.475a.5.5 0 1 0-.998.05c.083 1.672-.004 2.81-.238 3.374c-.183.442-.544.634-1.202.597a.5.5 0 0 0-.058.998c1.058.06 1.826-.349 2.184-1.211c.307-.74.403-1.996.312-3.808Z"/><path d="M6.5 2.5a.5.5 0 0 1 1 0v9a.5.5 0 0 1-1 0v-9Z"/><path fill-rule="evenodd" d="M5.5 12.5v5a1.5 1.5 0 0 0 3 0v-5a1.5 1.5 0 0 0-3 0ZM7 18a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 1 0v5a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}