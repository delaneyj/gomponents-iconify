package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XCircleSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.344 2.343h-.001a8 8 0 0 1 11.314 11.314A8.002 8.002 0 0 1 .234 10.089a8 8 0 0 1 2.11-7.746Zm1.06 10.253a6.5 6.5 0 1 0 9.108-9.275a6.5 6.5 0 0 0-9.108 9.275ZM6.03 4.97L8 6.94l1.97-1.97a.749.749 0 0 1 1.275.326a.749.749 0 0 1-.215.734L9.06 8l1.97 1.97a.749.749 0 0 1-.326 1.275a.749.749 0 0 1-.734-.215L8 9.06l-1.97 1.97a.749.749 0 0 1-1.275-.326a.749.749 0 0 1 .215-.734L6.94 8L4.97 6.03a.751.751 0 0 1 .018-1.042a.751.751 0 0 1 1.042-.018Z"/>`),
		g.Group(children),
	)
}