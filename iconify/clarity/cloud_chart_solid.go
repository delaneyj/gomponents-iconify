package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 7v22a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2Zm-23.636 6.364c-3.515 3.514-4.465 8.263-2.121 10.606c2.343 2.344 7.091 1.394 10.606-2.121c3.515-3.514 4.465-8.263 2.122-10.606s-7.092-1.394-10.607 2.121ZM25 19a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}