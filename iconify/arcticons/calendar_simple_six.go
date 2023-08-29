package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="28.333" r="4.667"/><path d="M28.148 20.728C27.457 19.691 26.247 19 24.346 19H24a4.647 4.647 0 0 0-4.667 4.667v4.666"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}