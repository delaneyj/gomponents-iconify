package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseWifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 18.965c-6.775-6.42-15.881-8.96-24.5-7.617M38 25.799a19.705 19.705 0 0 0-9.5-5.284M10 25.799a19.814 19.814 0 0 1 4.36-3.299M16 32.314a11.261 11.261 0 0 1 5-2.91"/><path fill="currentColor" fill-rule="evenodd" d="M24 40a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 40L8 8M4 18.965a29.317 29.317 0 0 1 3.5-2.84"/></g>`),
		g.Group(children),
	)
}