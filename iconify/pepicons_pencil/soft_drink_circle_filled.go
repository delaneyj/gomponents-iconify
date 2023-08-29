package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftDrinkCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilSoftDrinkCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M16.411 7.5H9.589a2.001 2.001 0 0 0-1.985 2.242l1.22 10A2 2 0 0 0 10.81 21.5h4.385a2 2 0 0 0 1.985-1.758l1.217-10A2 2 0 0 0 16.411 7.5ZM9.468 8.507A1 1 0 0 1 9.59 8.5h6.822a1 1 0 0 1 .993 1.12l-1.218 10a1 1 0 0 1-.992.88H10.81a1 1 0 0 1-.992-.879l-1.22-10a1 1 0 0 1 .871-1.114Z" clip-rule="evenodd"/><path d="M11.978 17.647a.5.5 0 1 1-.956-.294l4-13a.5.5 0 1 1 .956.294l-4 13Z"/><path d="M8.5 13a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1h-9Zm6.879-8.015a.5.5 0 0 1 .242-.97l4 1a.5.5 0 0 1-.242.97l-4-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilSoftDrinkCircleFilled0)"/></g>`),
		g.Group(children),
	)
}