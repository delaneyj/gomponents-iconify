package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.094 4.781L7.688 6.22l9.78 9.78l-9.78 9.781l1.406 1.438L20.313 16zm7 0L14.687 6.22L24.47 16l-9.782 9.781l1.407 1.438L27.312 16z"/>`),
		g.Group(children),
	)
}