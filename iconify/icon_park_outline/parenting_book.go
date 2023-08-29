package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParentingBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 31v13H13.625C9.875 44 8 42 8 37V4h12"/><circle cx="33" cy="17" r="9" stroke="currentColor" stroke-width="4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22 18v-2m22 2v-2M33 8c-.167-1-1.2-3.2-4-4m4 4c.083-1 .6-3.2 2-4"/><circle cx="36" cy="16" r="2" fill="currentColor"/><circle cx="30" cy="16" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 36h30"/></g>`),
		g.Group(children),
	)
}