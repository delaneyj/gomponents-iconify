package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Banana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M16 32c-2.8-.933-6.035-3.173-7.476-5.425a1.716 1.716 0 0 1-.245-1.043l.245-3.75c.053-.81 1.01-1.24 1.671-.77C12.33 22.529 16.228 25 19 25c7 0 13-2.5 16-8"/><path d="M26 43c10-2 19.242-12.485 16.867-23.059L41.999 16V8l-6-1c0 12.941-3 23-16 25c-5.976.92-11.705-.386-16.129-2.922L5 36c2 5 11 9 21 7Z"/></g>`),
		g.Group(children),
	)
}