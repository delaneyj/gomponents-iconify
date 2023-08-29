package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndexFinger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M12.566 26.182C10.856 27.354 10 29.294 10 32c0 4.06 4.975 11 9.462 11h11.48C35.331 43 38 39.15 38 36.06V23.01a3 3 0 0 0-3-3h-.01A2.99 2.99 0 0 0 32 23"/><path d="M13.981 28.445V8.005a2.998 2.998 0 0 1 3.006-2.997a3.014 3.014 0 0 1 3.006 3.015v15.569"/><path stroke-linejoin="round" d="M19.993 23.008v-3.992a3.016 3.016 0 0 1 6.03 0v3.992"/><path stroke-linejoin="round" d="M26 22.716v-2.713a3 3 0 0 1 6 0v3"/></g>`),
		g.Group(children),
	)
}