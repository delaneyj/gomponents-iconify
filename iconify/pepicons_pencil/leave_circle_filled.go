package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeaveCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilLeaveCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M18.346 10.116a.5.5 0 0 1 .705.064l2.083 2.5a.5.5 0 0 1-.768.64l-2.084-2.5a.5.5 0 0 1 .064-.704Z"/><path d="M18.346 15.884a.5.5 0 0 1-.064-.704l2.084-2.5a.5.5 0 1 1 .768.64l-2.083 2.5a.5.5 0 0 1-.704.064Z"/><path d="M20.5 13a.5.5 0 0 1-.5.5h-7.5a.5.5 0 0 1 0-1H20a.5.5 0 0 1 .5.5Zm-14-7a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1H7a.5.5 0 0 1-.5-.5Zm0 14a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1H7a.5.5 0 0 1-.5-.5Z"/><path d="M16 5.5a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5Zm0 10a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 .5-.5Zm-9-10a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilLeaveCircleFilled0)"/></g>`),
		g.Group(children),
	)
}