package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tubi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.365 20.384v4.962a3.007 3.007 0 0 0 3.007 3.007h0a3.007 3.007 0 0 0 3.008-3.007v-4.962"/><circle cx="35.447" cy="16.699" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.447 20.384v7.969M11.5 17.902v7.444a3.007 3.007 0 0 0 3.007 3.007h.408M11.5 20.384h3.039m11.801 3.007a3.007 3.007 0 0 1 3.008-3.007h0a3.007 3.007 0 0 1 3.008 3.007v1.955a3.007 3.007 0 0 1-3.008 3.007h0a3.007 3.007 0 0 1-3.007-3.007m0 0v-9.023"/>`),
		g.Group(children),
	)
}