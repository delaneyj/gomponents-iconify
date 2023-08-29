package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftDrinkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.411 4.5H6.589a2.001 2.001 0 0 0-1.985 2.242l1.22 10A2 2 0 0 0 7.81 18.5h4.385a2 2 0 0 0 1.985-1.758l1.217-10A2 2 0 0 0 13.411 4.5ZM6.468 5.507A1 1 0 0 1 6.59 5.5h6.822a1 1 0 0 1 .993 1.12l-1.218 10a1 1 0 0 1-.992.88H7.809a1 1 0 0 1-.992-.879l-1.22-10a1 1 0 0 1 .871-1.114Z" clip-rule="evenodd"/><path d="M8.978 14.647a.5.5 0 1 1-.956-.294l4-13a.5.5 0 1 1 .956.294l-4 13Z"/><path d="M5.5 10a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1h-9Zm6.879-8.015a.5.5 0 0 1 .242-.97l4 1a.5.5 0 0 1-.242.97l-4-1Z"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}