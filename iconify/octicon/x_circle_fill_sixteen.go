package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XCircleFillSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.343 13.657A8 8 0 1 1 13.658 2.343A8 8 0 0 1 2.343 13.657ZM6.03 4.97a.751.751 0 0 0-1.042.018a.751.751 0 0 0-.018 1.042L6.94 8L4.97 9.97a.749.749 0 0 0 .326 1.275a.749.749 0 0 0 .734-.215L8 9.06l1.97 1.97a.749.749 0 0 0 1.275-.326a.749.749 0 0 0-.215-.734L9.06 8l1.97-1.97a.749.749 0 0 0-.326-1.275a.749.749 0 0 0-.734.215L8 6.94Z"/>`),
		g.Group(children),
	)
}