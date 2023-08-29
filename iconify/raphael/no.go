package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func No(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3.667C9.19 3.667 3.667 9.187 3.667 16S9.19 28.333 16 28.333c6.812 0 12.333-5.52 12.333-12.333S22.813 3.667 16 3.667zm0 3c1.85 0 3.572.548 5.024 1.48L8.147 21.024A9.263 9.263 0 0 1 6.667 16c0-5.146 4.187-9.333 9.333-9.333zm0 18.666a9.271 9.271 0 0 1-5.024-1.48l12.876-12.877A9.263 9.263 0 0 1 25.332 16c0 5.146-4.186 9.333-9.332 9.333z"/>`),
		g.Group(children),
	)
}