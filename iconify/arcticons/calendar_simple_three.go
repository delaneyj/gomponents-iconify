package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.137 26a3.51 3.51 0 0 1 3.5 3.5h0a3.51 3.51 0 0 1-3.5 3.5h-1.4c-2.45 0-3.324-.35-4.374-1.225m0-11.55c1.05-.875 2.1-1.225 4.375-1.225h1.4a3.51 3.51 0 0 1 3.5 3.5h0a3.51 3.51 0 0 1-3.5 3.5h-3.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}