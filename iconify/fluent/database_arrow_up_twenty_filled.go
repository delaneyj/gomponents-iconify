package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseArrowUpTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 8c3.314 0 6-1.343 6-3s-2.686-3-6-3s-6 1.343-6 3s2.686 3 6 3Zm4.69.016c.47-.235.927-.534 1.31-.896v2.087a5.5 5.5 0 0 0-5.745 8.79L10 18c-3.314 0-6-1.343-6-3V7.12c.383.362.84.661 1.31.896C6.562 8.642 8.222 9 10 9c1.778 0 3.438-.358 4.69-.984ZM10 14.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Zm4.854-2.353l-.003-.003a.499.499 0 0 0-.348-.144h-.006a.5.5 0 0 0-.35.146l-2 2a.5.5 0 0 0 .707.708L14 13.707V16.5a.5.5 0 0 0 1 0v-2.793l1.146 1.147a.5.5 0 0 0 .708-.708l-2-2Z"/>`),
		g.Group(children),
	)
}