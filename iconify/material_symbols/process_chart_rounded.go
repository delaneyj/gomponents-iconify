package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProcessChartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.4 18q-.375-.2-.512-.588t.062-.762l5.1-10.2q.2-.375.588-.512T8.4 6q.375.2.513.588t-.063.762l-5.1 10.2q-.2.375-.588.513T2.4 18ZM9 18q-.375-.2-.513-.588t.063-.762l5.1-10.2q.2-.375.588-.512T15 6q.375.2.513.588t-.063.762l-5.1 10.2q-.2.375-.587.513T9 18Zm6.6 0q-.375-.2-.512-.588t.062-.762l5.1-10.2q.2-.375.588-.512T21.6 6q.375.2.513.588t-.063.762l-5.1 10.2q-.2.375-.588.513T15.6 18Z"/>`),
		g.Group(children),
	)
}