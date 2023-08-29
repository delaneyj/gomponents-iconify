package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M16.364 3.636a.75.75 0 0 0-1.06 1.06a7.5 7.5 0 0 1 0 10.607a.75.75 0 0 0 1.06 1.061a9 9 0 0 0 0-12.728ZM4.697 4.697a.75.75 0 0 0-1.061-1.06a9 9 0 0 0 0 12.727a.75.75 0 1 0 1.06-1.06a7.5 7.5 0 0 1 0-10.607Z"/><path d="M12.475 6.465a.75.75 0 0 1 1.06 0a5 5 0 0 1 0 7.07a.75.75 0 1 1-1.06-1.06a3.5 3.5 0 0 0 0-4.95a.75.75 0 0 1 0-1.06Zm-4.95 0a.75.75 0 0 1 0 1.06a3.5 3.5 0 0 0 0 4.95a.75.75 0 0 1-1.06 1.06a5 5 0 0 1 0-7.07a.75.75 0 0 1 1.06 0ZM11 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		g.Group(children),
	)
}