package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linkeye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.719 14.719L5.44 24l9.28 9.281L24 24ZM24 24l9.281 9.281L42.561 24l-9.28-9.281Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.304 24l2.415-2.415L17.135 24l-2.416 2.415Zm18.561 0l2.415-2.415L35.696 24l-2.416 2.415Z"/>`),
		g.Group(children),
	)
}