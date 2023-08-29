package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M44 39.82H4l32.54-18H44v18ZM3.857 19.201L7.09 20.8l17.321-10l.232-3.598"/><path fill="currentColor" d="M12.33 24.88a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm11.12-6a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`),
		g.Group(children),
	)
}