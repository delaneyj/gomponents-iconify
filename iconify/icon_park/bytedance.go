package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bytedance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M5 7L10 9V37L5 39V7Z"/><path d="M16 23L21 25V37L16 39V23Z"/><path d="M27 21L32 19V35L27 33V21Z"/><path d="M38 9L43 11V37L38 39V9Z"/></g>`),
		g.Group(children),
	)
}