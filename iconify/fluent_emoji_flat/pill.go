package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#FCD53F" d="m21.415 21.405l-6.39 6.39a7.629 7.629 0 0 1-10.79 0a7.629 7.629 0 0 1 0-10.79l6.39-6.39l8.642 2.585l2.148 8.205Z"/><path fill="#F8312F" d="m10.635 10.625l6.39-6.39a7.629 7.629 0 0 1 10.79 0a7.629 7.629 0 0 1 0 10.79l-6.39 6.39l-10.79-10.79Z"/><path fill="#F4F4F4" d="M26 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`),
		g.Group(children),
	)
}