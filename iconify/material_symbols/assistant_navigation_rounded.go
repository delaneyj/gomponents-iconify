package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssistantNavigationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm-3.3-5.375l3.075-1.35q.125-.05.238-.05t.237.05l3.05 1.35q.35.15.625-.113t.125-.612l-3.525-8.6q-.15-.35-.525-.35t-.525.35L7.95 15.9q-.15.35.125.613t.625.112Z"/>`),
		g.Group(children),
	)
}