package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveYCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilMoveYCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.616 9.653a.5.5 0 0 1 .064-.704l4-3.333a.5.5 0 1 1 .64.768l-4 3.333a.5.5 0 0 1-.704-.064Z"/><path d="M17.384 9.653a.5.5 0 0 1-.704.064l-4-3.333a.5.5 0 1 1 .64-.768l4 3.333a.5.5 0 0 1 .064.704ZM13 7.5a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-1 0V8a.5.5 0 0 1 .5-.5Zm4.384 8.846a.5.5 0 0 1-.064.705l-4 3.333a.5.5 0 0 1-.64-.768l4-3.334a.5.5 0 0 1 .704.064Z"/><path d="M8.616 16.346a.5.5 0 0 1 .704-.064l4 3.334a.5.5 0 0 1-.64.768l-4-3.333a.5.5 0 0 1-.064-.704Z"/><path d="M13 20a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 1 0v13a.5.5 0 0 1-.5.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilMoveYCircleFilled0)"/></g>`),
		g.Group(children),
	)
}