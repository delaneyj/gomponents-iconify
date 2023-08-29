package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareCircleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7.205 7.562a.75.75 0 0 0-.993-1.124A8.73 8.73 0 0 0 3.25 13a.75.75 0 0 0 1.5 0a7.23 7.23 0 0 1 2.455-5.438Zm10.583-1.124a.75.75 0 0 0-.993 1.124A7.23 7.23 0 0 1 19.25 13a.75.75 0 0 0 1.5 0a8.73 8.73 0 0 0-2.962-6.562Zm-7.601 13.584a.75.75 0 1 0-.374 1.452a8.773 8.773 0 0 0 4.374 0a.75.75 0 1 0-.374-1.452A7.267 7.267 0 0 1 12 20.25a7.31 7.31 0 0 1-1.813-.228Z" opacity=".5"/><path d="M9 6a3 3 0 1 0 6 0a3 3 0 0 0-6 0ZM2.5 18a3 3 0 1 0 6 0a3 3 0 0 0-6 0Zm16 3a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/></g>`),
		g.Group(children),
	)
}