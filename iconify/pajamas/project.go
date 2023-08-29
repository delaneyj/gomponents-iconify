package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Project(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m9.5 14.5l-6-2.5V4l6-2.5v13Zm-6.885-1.244A1 1 0 0 1 2 12.333V3.667a1 1 0 0 1 .615-.923L8.923.115A1.5 1.5 0 0 1 11 1.5V2h1.25c.966 0 1.75.783 1.75 1.75v8.5A1.75 1.75 0 0 1 12.25 14H11v.5a1.5 1.5 0 0 1-2.077 1.385l-6.308-2.629ZM11 12.5h1.25a.25.25 0 0 0 .25-.25v-8.5a.25.25 0 0 0-.25-.25H11v9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}