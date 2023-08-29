package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChineseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" rx="3"/><path d="M21 19.939c-.4-1.164-1.597-3.202-3.992-2.91c-2.396.29-4.193 3.492-3.993 7.567c.2 4.075 2.396 6.404 4.392 6.404C19.802 31 21 28.206 21 28.206M26 31V19m0 12v-6.5a4.5 4.5 0 0 1 4.5-4.5v0a4.5 4.5 0 0 1 4.5 4.5V31"/></g>`),
		g.Group(children),
	)
}