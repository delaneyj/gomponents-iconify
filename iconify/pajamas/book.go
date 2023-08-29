package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 2a2 2 0 0 1 2-2H13a.75.75 0 0 1 .75.75v10.5A.75.75 0 0 1 13 12H7.5v1.5H13a.75.75 0 0 1 0 1.5H7.5v1H6v-1H4.5a2.25 2.25 0 0 1-2.25-2.25V2ZM6 13.5V12H4.5a.75.75 0 0 0 0 1.5H6Zm6.25-12v9H4.5c-.263 0-.515.045-.75.128V2a.5.5 0 0 1 .5-.5h8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}