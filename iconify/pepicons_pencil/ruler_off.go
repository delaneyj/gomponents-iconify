package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2.31 13.626a.5.5 0 0 0 0 .708l3.536 3.535a.5.5 0 0 0 .707 0L17.867 6.555a.5.5 0 0 0 0-.707l-3.536-3.535a.5.5 0 0 0-.707 0l-1.06 1.06l1.709 1.71a.5.5 0 1 1-.708.706L11.857 4.08l-1.415 1.415l.884.884a.5.5 0 0 1-.707.707l-.884-.884l-1.414 1.414l1.709 1.709a.5.5 0 1 1-.707.707L7.614 8.323L6.2 9.737l.884.884a.5.5 0 1 1-.707.707l-.884-.884l-1.415 1.415l1.71 1.709a.5.5 0 1 1-.708.707l-1.709-1.71l-1.06 1.061Zm-.706 1.415a1.5 1.5 0 0 1 0-2.122L12.917 1.606a1.5 1.5 0 0 1 2.122 0l3.535 3.535a1.5 1.5 0 0 1 0 2.121L7.26 18.576a1.5 1.5 0 0 1-2.12 0l-3.536-3.535Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}