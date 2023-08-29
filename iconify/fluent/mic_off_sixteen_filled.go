package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10.809 11.516l3.337 3.338a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708L5.5 6.207V8a2.5 2.5 0 0 0 3.879 2.086l.717.717A3.5 3.5 0 0 1 4.5 8a.5.5 0 1 0-1 0a4.5 4.5 0 0 0 4 4.473V13.5a.5.5 0 1 0 1 0v-1.027a4.48 4.48 0 0 0 2.309-.957ZM12.06 9.94l-.764-.764c.132-.367.203-.763.203-1.176a.5.5 0 0 1 1 0c0 .695-.157 1.353-.439 1.94Zm-1.586-1.586L5.682 3.56a2.5 2.5 0 0 1 4.818.94V8c0 .12-.008.238-.025.354Z"/>`),
		g.Group(children),
	)
}