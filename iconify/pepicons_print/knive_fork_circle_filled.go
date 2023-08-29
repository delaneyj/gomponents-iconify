package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KniveForkCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill-rule="evenodd" d="M17.28 4.051C15.12 5.115 14 6.611 14 8.506c0 1.82 1.037 3.051 3 3.631V14a.5.5 0 0 0 1 0V4.5a.5.5 0 0 0-.72-.449ZM15 8.506c0-1.268.65-2.314 2-3.158v5.74c-1.364-.477-2-1.321-2-2.582Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16 15.5v5a1.5 1.5 0 0 0 3 0v-5a1.5 1.5 0 0 0-3 0Zm1.5 5.5a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 1 0v5a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path d="M7.532 5.475a.5.5 0 1 1 1 .05c-.084 1.672.003 2.81.237 3.374c.183.442.544.634 1.202.597a.5.5 0 0 1 .058.998c-1.058.06-1.826-.349-2.184-1.211c-.307-.74-.403-1.996-.313-3.808Z"/><path d="M12.467 5.475a.5.5 0 1 0-.998.05c.083 1.672-.004 2.81-.238 3.374c-.183.442-.544.634-1.202.597a.5.5 0 0 0-.058.998c1.058.06 1.826-.349 2.184-1.211c.307-.74.403-1.996.312-3.808Z"/><path d="M9.5 5.5a.5.5 0 0 1 1 0v9a.5.5 0 0 1-1 0v-9Z"/><path fill-rule="evenodd" d="M8.5 15.5v5a1.5 1.5 0 0 0 3 0v-5a1.5 1.5 0 0 0-3 0ZM10 21a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 1 0v5a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}