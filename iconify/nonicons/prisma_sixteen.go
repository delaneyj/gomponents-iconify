package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrismaSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.214 2.043a.25.25 0 0 0-.428 0l-5.564 9.141a.25.25 0 0 0 .109.357l5.564 2.583a.25.25 0 0 0 .21 0l5.564-2.583a.25.25 0 0 0 .109-.357l-5.564-9.14Zm-1.709-.78a1.75 1.75 0 0 1 2.99 0l5.564 9.141a1.75 1.75 0 0 1-.758 2.497l-5.564 2.584a1.75 1.75 0 0 1-1.474 0L1.699 12.9a1.75 1.75 0 0 1-.758-2.497l5.564-9.14Z"/>`),
		g.Group(children),
	)
}