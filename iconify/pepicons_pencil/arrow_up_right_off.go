package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M7.852 5.952a.5.5 0 0 1 .543-.453l5.186.472a.5.5 0 0 1-.09.996l-5.186-.472a.5.5 0 0 1-.453-.543Z"/><path d="M14.052 12.152a.5.5 0 0 1-.543-.453l-.472-5.185a.5.5 0 0 1 .996-.09l.472 5.185a.5.5 0 0 1-.453.543Z"/><path d="M13.354 6.65a.5.5 0 0 1 0 .708l-6.536 6.535a.5.5 0 0 1-.707-.707l6.535-6.536a.5.5 0 0 1 .707 0Z"/></g><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}