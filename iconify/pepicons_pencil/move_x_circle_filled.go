package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveXCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilMoveXCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M16.347 8.616a.5.5 0 0 1 .704.064l3.333 4a.5.5 0 0 1-.768.64l-3.333-4a.5.5 0 0 1 .064-.704Z"/><path d="M16.347 17.384a.5.5 0 0 1-.064-.704l3.333-4a.5.5 0 1 1 .768.64l-3.333 4a.5.5 0 0 1-.704.064Z"/><path d="M6.5 12.5h13a.5.5 0 1 1 0 1h-13a.5.5 0 0 1 0-1Z"/><path d="M9.653 17.384a.5.5 0 0 1-.704-.064l-3.333-4a.5.5 0 0 1 .768-.64l3.333 4a.5.5 0 0 1-.064.704Z"/><path d="M9.653 8.616a.5.5 0 0 1 .064.704l-3.333 4a.5.5 0 0 1-.768-.64l3.333-4a.5.5 0 0 1 .704-.064ZM7.5 13a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1H8a.5.5 0 0 1-.5-.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilMoveXCircleFilled0)"/></g>`),
		g.Group(children),
	)
}