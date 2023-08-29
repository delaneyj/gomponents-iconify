package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CimbBankPh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 24H12.798L4.5 7.404h30.702L43.5 24zm-30.702 0L4.5 40.596h30.702L43.5 24"/>`),
		g.Group(children),
	)
}