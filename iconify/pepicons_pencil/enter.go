package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.547 7.116a.5.5 0 0 1 .704.064l2.083 2.5a.5.5 0 0 1-.768.64l-2.083-2.5a.5.5 0 0 1 .064-.704Z"/><path d="M8.546 12.884a.5.5 0 0 1-.064-.704l2.084-2.5a.5.5 0 1 1 .768.64l-2.083 2.5a.5.5 0 0 1-.704.064Z"/><path d="M11 10a.5.5 0 0 1-.5.5H3a.5.5 0 0 1 0-1h7.5a.5.5 0 0 1 .5.5Zm5 7a.5.5 0 0 1-.5.5H6.3a.5.5 0 0 1 0-1h9.2a.5.5 0 0 1 .5.5Zm0-14a.5.5 0 0 1-.5.5H6.326a.5.5 0 1 1 0-1H15.5a.5.5 0 0 1 .5.5Z"/><path d="M6.25 17.5a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 1 0v4a.5.5 0 0 1-.5.5Zm0-10a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 1 0v4a.5.5 0 0 1-.5.5Zm9.35 10a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 1 0v14a.5.5 0 0 1-.5.5Z"/></g>`),
		g.Group(children),
	)
}