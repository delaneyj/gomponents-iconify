package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityFourteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1.923l-6.3 2.52l.743 1.857L7 6.077V8H1v14h22V8h-6V6.077l.557.223l.743-1.857l-6.3-2.52Zm3 3.354V20h-2v-5h-2v5H9V5.277l3-1.2l3 1.2ZM7 20H3V10h4v10Zm10 0V10h4v10h-4ZM13.004 6.998H11v2.004h2.004V6.998Z"/>`),
		g.Group(children),
	)
}