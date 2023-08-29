package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilEnterCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.547 10.116a.5.5 0 0 1 .704.064l2.083 2.5a.5.5 0 0 1-.768.64l-2.083-2.5a.5.5 0 0 1 .064-.704Z"/><path d="M11.546 15.884a.5.5 0 0 1-.064-.704l2.084-2.5a.5.5 0 1 1 .768.64l-2.083 2.5a.5.5 0 0 1-.704.064Z"/><path d="M14 13a.5.5 0 0 1-.5.5H6a.5.5 0 0 1 0-1h7.5a.5.5 0 0 1 .5.5Zm5 7a.5.5 0 0 1-.5.5H9.3a.5.5 0 0 1 0-1h9.2a.5.5 0 0 1 .5.5Zm0-14a.5.5 0 0 1-.5.5H9.326a.5.5 0 1 1 0-1H18.5a.5.5 0 0 1 .5.5Z"/><path d="M9.25 20.5a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 1 0v4a.5.5 0 0 1-.5.5Zm0-10a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 1 0v4a.5.5 0 0 1-.5.5Zm9.35 10a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 1 0v14a.5.5 0 0 1-.5.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilEnterCircleFilled0)"/></g>`),
		g.Group(children),
	)
}