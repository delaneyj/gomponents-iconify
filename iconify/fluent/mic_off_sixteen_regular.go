package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOffSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10.809 11.516l3.337 3.338a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708L5.5 6.207V8a2.5 2.5 0 0 0 3.879 2.086l.717.717A3.5 3.5 0 0 1 4.5 8a.5.5 0 1 0-1 0a4.5 4.5 0 0 0 4 4.473V13.5a.5.5 0 1 0 1 0v-1.027a4.48 4.48 0 0 0 2.309-.957ZM8.647 9.354A1.5 1.5 0 0 1 6.5 8v-.793l2.147 2.147ZM9.5 7.379V4.5a1.5 1.5 0 0 0-2.996-.117l-.822-.822A2.5 2.5 0 0 1 10.5 4.5V8c0 .12-.008.238-.025.354L9.5 7.379Zm2.561 2.561l-.764-.764c.132-.367.203-.763.203-1.176a.5.5 0 0 1 1 0c0 .695-.157 1.353-.439 1.94Z"/>`),
		g.Group(children),
	)
}