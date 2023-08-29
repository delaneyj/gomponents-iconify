package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextGrammarOptionsTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M13.78 16.216a.75.75 0 0 1 0 1.06l-4.5 4.5a.75.75 0 1 1-1.06-1.06l4.5-4.5a.75.75 0 0 1 1.06 0zM10.522 17l-2 2H3a1 1 0 0 1-.117-1.993L3 17h7.522zM16.5 8a.75.75 0 0 1 .744.658l.14 1.13a3.25 3.25 0 0 0 2.828 2.829l1.13.139a.75.75 0 0 1 0 1.488l-1.13.14a3.25 3.25 0 0 0-2.829 2.828l-.139 1.13a.75.75 0 0 1-1.488 0l-.14-1.13a3.25 3.25 0 0 0-2.828-2.829l-1.13-.139a.75.75 0 0 1 0-1.488l1.13-.14a3.25 3.25 0 0 0 2.829-2.828l.139-1.13A.75.75 0 0 1 16.5 8zm-6.427 5a1.75 1.75 0 0 0 .65 1.918l.125.082H3a1 1 0 0 1-.23-1.974l.113-.02L3 13h7.073zM13 9a1 1 0 0 1 .117 1.993L13 11H3a1 1 0 0 1-.117-1.993L3 9h10zm8-4a1 1 0 0 1 .23 1.974l-.113.02L21 7H3a1 1 0 0 1-.23-1.974l.113-.02L3 5h18z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}