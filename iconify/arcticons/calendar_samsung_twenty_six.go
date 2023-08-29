package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSamsungTwentySix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.72 21.667c0-2.593 2.073-4.667 4.493-4.667s4.667 2.074 4.667 4.667c0 1.21-.519 2.42-1.383 3.284C19.596 26.506 13.72 31 13.72 31h9.16"/><circle cx="29.614" cy="26.333" r="4.667"/><path d="M33.762 18.728C33.071 17.691 31.861 17 29.96 17h-.346a4.647 4.647 0 0 0-4.667 4.667v4.666"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 5.5h-4v1a3.5 3.5 0 1 1-7 0v-1h-8v1a3.5 3.5 0 1 1-7 0v-1h-3a4 4 0 0 0-4 4v29a4 4 0 0 0 4 4h29a4 4 0 0 0 4-4v-29a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}