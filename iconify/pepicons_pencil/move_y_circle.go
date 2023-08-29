package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveYCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.616 9.653a.5.5 0 0 1 .064-.704l4-3.333a.5.5 0 1 1 .64.768l-4 3.333a.5.5 0 0 1-.704-.064Z"/><path d="M17.384 9.653a.5.5 0 0 1-.704.064l-4-3.333a.5.5 0 1 1 .64-.768l4 3.333a.5.5 0 0 1 .064.704ZM13 7.5a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-1 0V8a.5.5 0 0 1 .5-.5Zm4.384 8.846a.5.5 0 0 1-.064.705l-4 3.333a.5.5 0 0 1-.64-.768l4-3.334a.5.5 0 0 1 .704.064Z"/><path d="M8.616 16.346a.5.5 0 0 1 .704-.064l4 3.334a.5.5 0 0 1-.64.768l-4-3.333a.5.5 0 0 1-.064-.704Z"/><path d="M13 20a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 1 0v13a.5.5 0 0 1-.5.5Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}