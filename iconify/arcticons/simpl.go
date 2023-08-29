package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simpl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.923 14.324c6.675 6.675 6.675 17.496 0 24.17s-17.495 6.675-24.17 0m2.323-4.818c-6.674-6.675-6.674-17.496 0-24.17s17.496-6.675 24.17 0"/>`),
		g.Group(children),
	)
}