package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveXOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M13.347 5.616a.5.5 0 0 1 .704.064l3.333 4a.5.5 0 0 1-.768.64l-3.333-4a.5.5 0 0 1 .064-.704Z"/><path d="M13.347 14.384a.5.5 0 0 1-.064-.704l3.333-4a.5.5 0 1 1 .768.64l-3.333 4a.5.5 0 0 1-.704.064Z"/><path d="M3.5 9.5h13a.5.5 0 1 1 0 1h-13a.5.5 0 0 1 0-1Z"/><path d="M6.653 14.384a.5.5 0 0 1-.704-.064l-3.333-4a.5.5 0 0 1 .768-.64l3.333 4a.5.5 0 0 1-.064.704Z"/><path d="M6.653 5.616a.5.5 0 0 1 .064.704l-3.333 4a.5.5 0 0 1-.768-.64l3.333-4a.5.5 0 0 1 .704-.064ZM4.5 10a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Z"/></g><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}