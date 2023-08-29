package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 11.03v4h-2v-4H8v6h4v4h2v-10h-2zm12.19 0H22l-3 4.39v-4.39h-2v10h2V18.3l.91-1.33L22 21.03h2.19l-2.99-5.62l2.99-4.38z"/><path fill="currentColor" d="M28 26H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h24a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2ZM4 8v16h24V8Z"/>`),
		g.Group(children),
	)
}