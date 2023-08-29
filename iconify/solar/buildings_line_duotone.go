package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingsLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M22 22H2"/><path d="M21 22V6c0-1.886 0-2.828-.586-3.414C19.828 2 18.886 2 17 2h-2c-1.886 0-2.828 0-3.414.586c-.472.471-.564 1.174-.582 2.414" opacity=".5"/><path d="M15 22V9c0-1.886 0-2.828-.586-3.414C13.828 5 12.886 5 11 5H7c-1.886 0-2.828 0-3.414.586C3 6.172 3 7.114 3 9v13"/><path stroke-linecap="round" d="M9 22v-3"/><path stroke-linecap="round" d="M6 8h6m-6 3h6m-6 3h6" opacity=".5"/></g>`),
		g.Group(children),
	)
}