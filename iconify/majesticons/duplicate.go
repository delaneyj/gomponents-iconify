package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M16 6a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v6a4 4 0 0 0 4 4h1v-4.5c0-1.763.746-2.913 1.708-3.606C9.644 7.22 10.753 7 11.5 7H16V6zm.5 2h-5c-.587 0-1.478.18-2.208.706C8.588 9.213 8 10.063 8 11.5V18a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-6a4 4 0 0 0-4-4h-1.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}