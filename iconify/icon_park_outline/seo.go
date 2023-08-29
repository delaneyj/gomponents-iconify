package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="40" height="32" x="4" y="8" rx="1.633"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 18.948c-2-2.948-5.502-1.01-5.251 2.02C11 24 15 24 15.249 27.032C15.5 30.062 12 32 10 29.051M26 18h-4v13h4m-4-6h4"/><rect width="6" height="13" x="32" y="18" stroke-linecap="round" stroke-linejoin="round" rx="3"/></g>`),
		g.Group(children),
	)
}