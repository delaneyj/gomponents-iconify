package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 96H24l104-64Z" opacity=".2"/><path d="M24 104h24v64H32a8 8 0 0 0 0 16h192a8 8 0 0 0 0-16h-16v-64h24a8 8 0 0 0 4.19-14.81l-104-64a8 8 0 0 0-8.38 0l-104 64A8 8 0 0 0 24 104Zm40 0h32v64H64Zm80 0v64h-32v-64Zm48 64h-32v-64h32ZM128 41.39L203.74 88H52.26ZM248 208a8 8 0 0 1-8 8H16a8 8 0 0 1 0-16h224a8 8 0 0 1 8 8Z"/></g>`),
		g.Group(children),
	)
}