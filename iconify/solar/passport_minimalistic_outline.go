package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassportMinimalisticOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.25 14a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0ZM12 11.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5Z"/><path d="M18.75 4.306a2.75 2.75 0 0 0-3.139-2.722L4.814 3.126A1.821 1.821 0 0 0 3.25 5v14A3.75 3.75 0 0 0 7 22.75h10A3.75 3.75 0 0 0 20.75 19V9a3.75 3.75 0 0 0-2-3.317V4.306ZM5.062 6.75H17.01A2.25 2.25 0 0 1 19.25 9v10A2.25 2.25 0 0 1 17 21.25H7A2.25 2.25 0 0 1 4.75 19V6.75h.312Zm11.95-1.5h.238v-.944a1.25 1.25 0 0 0-1.427-1.237L5.026 4.61a.321.321 0 0 0 .041.639h11.945Z"/></g>`),
		g.Group(children),
	)
}