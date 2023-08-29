package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosChevronTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M13.97 4.72a.75.75 0 0 0 0 1.06L20.19 12l-6.22 6.22a.75.75 0 1 0 1.06 1.06l6.75-6.75a.75.75 0 0 0 0-1.06l-6.75-6.75a.75.75 0 0 0-1.06 0z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}