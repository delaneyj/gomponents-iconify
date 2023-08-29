package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Athmvil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.129 36.279l10.45 7.221V29.057l-10.45 7.222zm10.451 0h17.094a6.077 6.077 0 0 0 6.076-6.077v-11.03m2.121-7.451L31.421 4.5v14.443l10.45-7.222zm-10.451 0H14.327a6.077 6.077 0 0 0-6.076 6.077v11.03"/>`),
		g.Group(children),
	)
}