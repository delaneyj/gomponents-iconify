package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m14.889 20.925l3.5-1.925m0 0v14"/><circle cx="26.444" cy="28.333" r="4.667"/><path d="M30.593 20.728C29.9 19.691 28.69 19 26.79 19h-.346a4.647 4.647 0 0 0-4.666 4.667v4.666"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}