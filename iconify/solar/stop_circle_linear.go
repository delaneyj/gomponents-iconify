package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopCircleLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path d="M8 12c0-1.886 0-2.828.586-3.414C9.172 8 10.114 8 12 8c1.886 0 2.828 0 3.414.586C16 9.172 16 10.114 16 12c0 1.886 0 2.828-.586 3.414C14.828 16 13.886 16 12 16c-1.886 0-2.828 0-3.414-.586C8 14.828 8 13.886 8 12Z"/></g>`),
		g.Group(children),
	)
}