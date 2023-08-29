package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeSquaresLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 12H6c-1.886 0-2.828 0-3.414.586C2 13.172 2 14.114 2 16v2c0 1.886 0 2.828.586 3.414C3.172 22 4.114 22 6 22h2c1.886 0 2.828 0 3.414-.586C12 20.828 12 19.886 12 18v-1"/><path d="M12 7h-1c-1.886 0-2.828 0-3.414.586C7 8.172 7 9.114 7 11v2c0 1.886 0 2.828.586 3.414C8.172 17 9.114 17 11 17h2c1.886 0 2.828 0 3.414-.586C17 15.828 17 14.886 17 13v-1"/><path d="M12 6c0-1.886 0-2.828.586-3.414C13.172 2 14.114 2 16 2h2c1.886 0 2.828 0 3.414.586C22 3.172 22 4.114 22 6v2c0 1.886 0 2.828-.586 3.414C20.828 12 19.886 12 18 12h-2c-1.886 0-2.828 0-3.414-.586C12 10.828 12 9.886 12 8V6Z"/></g>`),
		g.Group(children),
	)
}