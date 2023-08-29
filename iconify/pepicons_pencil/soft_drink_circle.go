package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftDrinkCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M16.411 7.5H9.589a2.001 2.001 0 0 0-1.985 2.242l1.22 10A2 2 0 0 0 10.81 21.5h4.385a2 2 0 0 0 1.985-1.758l1.217-10A2 2 0 0 0 16.411 7.5ZM9.468 8.507A1 1 0 0 1 9.59 8.5h6.822a1 1 0 0 1 .993 1.12l-1.218 10a1 1 0 0 1-.992.88H10.81a1 1 0 0 1-.992-.879l-1.22-10a1 1 0 0 1 .871-1.114Z" clip-rule="evenodd"/><path d="M11.978 17.647a.5.5 0 1 1-.956-.294l4-13a.5.5 0 1 1 .956.294l-4 13Z"/><path d="M8.5 13a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1h-9Zm6.879-8.015a.5.5 0 0 1 .242-.97l4 1a.5.5 0 0 1-.242.97l-4-1Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}