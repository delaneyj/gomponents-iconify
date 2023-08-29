package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoSecurityTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.25A3.25 3.25 0 0 1 5.25 3h8.5A3.25 3.25 0 0 1 17 6.25v6.5A3.25 3.25 0 0 1 13.75 16h-8.5A3.25 3.25 0 0 1 2 12.75v-6.5Zm19.62-2.653a.75.75 0 0 1 .38.653v10.5a.75.75 0 0 1-1.136.643L18 13.675v-8.35l2.864-1.718a.75.75 0 0 1 .755-.01ZM8.135 17a2.501 2.501 0 0 1-2.386 1.75h-2a.75.75 0 0 0-.75.75v1.6a.9.9 0 0 0 .9.9h1.944a6.158 6.158 0 0 0 6.048-5H8.136Z"/>`),
		g.Group(children),
	)
}