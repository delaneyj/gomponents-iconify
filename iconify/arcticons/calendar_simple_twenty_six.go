package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleTwentySix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.72 23.667c0-2.593 2.073-4.667 4.493-4.667s4.667 2.074 4.667 4.667c0 1.21-.519 2.42-1.383 3.284C19.596 28.506 13.72 33 13.72 33h9.16"/><circle cx="29.614" cy="28.333" r="4.667"/><path d="M33.762 20.728C33.071 19.691 31.861 19 29.96 19h-.346a4.647 4.647 0 0 0-4.667 4.667v4.666"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}