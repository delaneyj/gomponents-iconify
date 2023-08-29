package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalSpacingLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M22 3H2m20 18H2"/><path d="M20 12c0-1.886 0-2.828-.586-3.414C18.828 8 17.886 8 16 8H8c-1.886 0-2.828 0-3.414.586C4 9.172 4 10.114 4 12c0 1.886 0 2.828.586 3.414C5.172 16 6.114 16 8 16h8c1.886 0 2.828 0 3.414-.586C20 14.828 20 13.886 20 12Z"/></g>`),
		g.Group(children),
	)
}