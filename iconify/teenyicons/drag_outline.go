package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5 7.5l.137-.48a.5.5 0 0 0-.618.617L7.5 7.5Zm2 7l-.48.137a.5.5 0 0 0 .94.06L9.5 14.5Zm5-5l.197.46a.5.5 0 0 0-.06-.94l-.137.48ZM11 11l-.197-.46a.5.5 0 0 0-.263.263L11 11ZM3.5 3.5V3H3v.5h.5Zm10 0h.5V3h-.5v.5Zm-10 10H3v.5h.5v-.5ZM0 1h1V0H0v1Zm4 0h1V0H4v1Zm4 0h1V0H8v1ZM0 5h1V4H0v1Zm0 4h1V8H0v1Zm7.02-1.363l2 7l.96-.274l-2-7l-.96.274Zm7.617 1.382l-7-2l-.274.962l7 2l.274-.962ZM9.96 14.697l1.5-3.5l-.92-.394l-1.5 3.5l.92.394Zm1.237-3.237l3.5-1.5l-.394-.92l-3.5 1.5l.394.92ZM3.5 4h10V3h-10v1Zm9.5-.5V7h1V3.5h-1Zm-10 0v10h1v-10H3ZM3.5 14H7v-1H3.5v1Z"/>`),
		g.Group(children),
	)
}