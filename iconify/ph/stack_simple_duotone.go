package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m240 104l-112 64l-112-64l112-64Z" opacity=".2"/><path d="m12 111l112 64a8 8 0 0 0 7.94 0l112-64a8 8 0 0 0 0-13.9l-112-64a8 8 0 0 0-7.94 0l-112 64a8 8 0 0 0 0 13.9Zm116-61.79L223.87 104L128 158.79L32.13 104ZM247 140a8 8 0 0 1-3 11l-112 64a8 8 0 0 1-7.94 0L12 151a8 8 0 1 1 8-13.95l108 61.74l108-61.74a8 8 0 0 1 11 2.95Z"/></g>`),
		g.Group(children),
	)
}