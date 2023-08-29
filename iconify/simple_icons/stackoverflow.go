package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stackoverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.725 0l-1.72 1.277l6.39 8.588l1.716-1.277L15.725 0zm-3.94 3.418l-1.369 1.644l8.225 6.85l1.369-1.644l-8.225-6.85zm-3.15 4.465l-.905 1.94l9.702 4.517l.904-1.94l-9.701-4.517zm-1.85 4.86l-.44 2.093l10.473 2.201l.44-2.092l-10.473-2.203zM1.89 15.47V24h19.19v-8.53h-2.133v6.397H4.021v-6.396H1.89zm4.265 2.133v2.13h10.66v-2.13H6.154Z"/>`),
		g.Group(children),
	)
}