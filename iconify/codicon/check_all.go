package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.62 3.596L7.815 12.81l-.728-.033L4 8.382l.754-.53l2.744 3.907L14.917 3l.703.596z"/><path d="m7.234 8.774l4.386-5.178L10.917 3l-4.23 4.994l.547.78zm-1.55.403l.548.78l-.547-.78zm-1.617 1.91l.547.78l-.799.943l-.728-.033L0 8.382l.754-.53l2.744 3.907l.57-.672z"/></g>`),
		g.Group(children),
	)
}