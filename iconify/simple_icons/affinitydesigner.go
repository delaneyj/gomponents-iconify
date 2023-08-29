package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Affinitydesigner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.44 0L0 18.083v5.197a.72.72 0 0 0 .713.72h10.023L5.7 15.277L14.52 0zm5.16 0l-4.86 8.418l3.718 6.439H24V.718A.72.72 0 0 0 23.28 0zm-5.4 9.353l-2.064 3.575a1.289 1.289 0 0 0 0 1.288c.23.4.656.64 1.117.64h4.125zm-3.122 6.44L11.816 24h11.471a.72.72 0 0 0 .713-.718v-7.49Z"/>`),
		g.Group(children),
	)
}