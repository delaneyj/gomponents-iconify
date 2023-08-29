package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.454 6.515a7 7 0 0 0 9.031 9.031l.766.766A7.972 7.972 0 0 1 12 17a7.966 7.966 0 0 1-5-1.754l.003 6a.75.75 0 0 0 1.181.612l3.817-2.687l3.818 2.687a.75.75 0 0 0 1.182-.613v-3.183l3.718 3.718a.75.75 0 0 0 1.061-1.06L3.28 2.22a.75.75 0 1 0-1.06 1.06l3.234 3.235ZM7.142 3.96l9.898 9.898A7 7 0 0 0 7.142 3.96Z"/>`),
		g.Group(children),
	)
}