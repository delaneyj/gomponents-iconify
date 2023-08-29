package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFootnoteGaNaTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M20.023 7.542a5.21 5.21 0 0 1-.218.154a1 1 0 1 1-1.11-1.664c.724-.483 1.241-1.244 1.35-1.745a1 1 0 0 1 1.978.213V11a1 1 0 1 1-2 0V7.542zM1 9a1 1 0 0 1 1-1h3.5a1 1 0 0 1 1 1c0 .991-.169 2.31-.789 3.475c-.643 1.209-1.768 2.234-3.556 2.513a1 1 0 1 1-.31-1.976c1.087-.17 1.713-.747 2.1-1.476c.25-.47.401-1.006.481-1.536H2a1 1 0 0 1-1-1zm6.5-3v12a1 1 0 0 0 2 0V6a1 1 0 0 0-2 0zm8 12V6a1 1 0 1 1 2 0v12a1 1 0 0 1-2 0zM13.478 7.21a1 1 0 1 0-1.956-.42l-1.5 7a1 1 0 0 0 1.142 1.196l3-.5a1 1 0 1 0-.328-1.972l-1.55.258l1.192-5.562z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}