package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCardBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 1.75a1.25 1.25 0 1 0 0 2.5v-2.5ZM5 3l1.246-.096A1.25 1.25 0 0 0 5 1.75V3Zm16 3l1.242.138A1.25 1.25 0 0 0 21 4.75V6ZM5.23 6l-1.246.096L5.231 6Zm13.109 9.119l.089 1.247l-.09-1.247Zm-10.355.74l-.089-1.248l.09 1.247ZM3 4.25h2v-2.5H3v2.5Zm5.073 12.855l10.355-.74l-.178-2.493l-10.355.74l.178 2.493Zm13.353-3.622l.816-7.345l-2.484-.276l-.816 7.345l2.484.276ZM3.754 3.096l.23 3l2.493-.192l-.23-3l-2.493.192Zm.23 3l.617 8.017l2.493-.192l-.617-8.017l-2.493.192ZM21 4.75H5.23v2.5H21v-2.5Zm-2.572 11.616a3.25 3.25 0 0 0 2.998-2.883l-2.484-.276a.75.75 0 0 1-.692.665l.178 2.494ZM7.895 14.61a.75.75 0 0 1-.801-.69l-2.493.192a3.25 3.25 0 0 0 3.472 2.992l-.178-2.494Z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3.75" d="M8.5 20.5h.01v.01H8.5zm9 0h.01v.01h-.01z"/></g>`),
		g.Group(children),
	)
}