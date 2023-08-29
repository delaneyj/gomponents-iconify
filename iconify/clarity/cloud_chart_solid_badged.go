package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudChartSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 12.34V29a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h18.57a7.44 7.44 0 0 0 6.74 8.46l1.01.02A7.453 7.453 0 0 0 34 12.34Zm-23.636 1.024c-3.515 3.514-4.465 8.263-2.121 10.606c2.343 2.344 7.091 1.394 10.606-2.121c3.515-3.514 4.465-8.263 2.122-10.606s-7.092-1.394-10.607 2.121ZM25 19a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}