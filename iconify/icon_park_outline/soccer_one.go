package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoccerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-miterlimit="2" stroke-width="4" d="M29 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path fill="currentColor" d="M19 46a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M24.38 18.91L20 27l7.31 6.02c.31.24.53.57.66.94L31 44M20 27l-5.49 7.64c-.54.79-1.56 1.09-2.44.72L4 32"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M42 24h-7c-.36 0-.72-.09-1.04-.27l-3.21-1.81c-4.01-2.35-8.4-3.96-12.97-4.78l-3.22-.57c-.45-.08-.92 0-1.32.22L7 20"/></g>`),
		g.Group(children),
	)
}